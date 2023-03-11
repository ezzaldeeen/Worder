package generator

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const samplePath = "./resources/sample.txt"

type TxtFileGenerator struct {
	Size        int
	Unit        string
	Count       int
	Destination string
}

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
		created, err := os.Create(g.Destination + fmt.Sprintf("/%s", strconv.Itoa(i)))
		if err != nil {
			return err
		}
		// todo: create unit converter
		err = write(created, body, g.Size)
		if err != nil {
			return err
		}
		created.Close()
	}

	return nil
}

func load(file *os.File) (string, error) {
	body, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read, %w", err)
	}
	return string(body), nil
}

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
