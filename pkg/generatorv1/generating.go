package generatorv1

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// todo: use viper to set the configuration
const samplePath = "./resources/sample.txt"

// TxtFileGenerator is a TXT file type generator
// it generates new files based on the given fields
// and uses the sample from Lorem Ipsum
type TxtFileGenerator struct {
	Size        int
	Unit        string
	Count       int
	Destination string
}

// Generate the only exposed method in the TxtFileGenerator
// it uses all the utils in order to generate new files
func (g TxtFileGenerator) Generate() error {
	file, err := os.Open(samplePath)
	if err != nil {
		return fmt.Errorf("unable to open, %w", err)
	}
	defer file.Close()

	body, err := load(file)
	if err != nil {
		return err
	}
	err = mkdir(g.Destination)
	if err != nil {
		return err
	}

	for i := 0; i < g.Count; i++ {
		cfile, err := os.Create(g.Destination + fmt.Sprintf("/%s", strconv.Itoa(i)))
		if err != nil {
			return err
		}

		size := convert(g.Size, g.Unit)
		err = write(cfile, body, size)
		if err != nil {
			return err
		}
		cfile.Close()
	}

	return nil
}

// load is for getting the content of file
func load(file *os.File) (string, error) {
	body, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read, %w", err)
	}
	return string(body), nil
}

// write on a file based on the given payload and size
func write(file *os.File, payload string, size int) error {
	for totalWrittenBytes := 0; totalWrittenBytes <= size; {
		writtenByte, err := io.WriteString(file, payload)
		if err != nil {
			return fmt.Errorf("unable to write, %w", err)
		}
		totalWrittenBytes += writtenByte
	}
	return nil
}

// mkdir make new directory based on the specified target
func mkdir(destination string) error {
	// make sure that the old directory has removed
	err := os.RemoveAll(destination)
	if err != nil {
		return fmt.Errorf("unable to remove. %w", err)
	}
	err = os.Mkdir(destination, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to make directory. %w", err)
	}
	return nil
}

// convert for unit (MB, KB) standardization into byte unit
func convert(size int, unit string) int {
	switch unit {
	case "mb":
		size *= 1000000
	case "kb":
		size *= 1000
	}
	return size
}
