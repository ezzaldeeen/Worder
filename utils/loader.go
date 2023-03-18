package utils

import (
	"bufio"
	"io"
)

// GetFileContent it's reading the file line-by-line
// and return the final combination all lines in the given file
func GetFileContent(reader *bufio.Reader) (string, error) {
	var lines string
loop:
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			switch err {
			case io.EOF:
				break loop
			default:
				return "", err
			}
		}
		lines += line
	}
	return lines, nil
}
