package counterv1

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type Counter struct {
	Path string
}

func (c Counter) Count() (int, error) {
	counter := 0
	files, _ := os.ReadDir(c.Path)
	for _, file := range files {
		filePath := path.Join("data", file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return 0, err
		}
		buf, _ := io.ReadAll(f)
		err = f.Close()
		if err != nil {
			return 0, err
		}

		content := string(buf)
		counter += len(strings.Fields(content))
	}
	return counter, nil
}

func load(file *os.File) (string, error) {
	body, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read, %w", err)
	}
	return string(body), nil
}
