package utils

import (
	"bufio"
	"strings"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	expected := "Fake dummy text\n"
	fakeReader := strings.NewReader(expected)
	actual, err := GetFileContent(bufio.NewReader(fakeReader))
	if err != nil {
		t.Errorf("Expected: %s, but Actual: %s", expected, actual)
	}
	if !strings.EqualFold(expected, actual) {
		t.Errorf("Expected: %s - Actual: %s", expected, actual)
	}
}
