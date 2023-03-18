package utils

import (
	"strconv"
	"strings"
	"worder/custom"
)

type unitTag string

const (
	KB unitTag = "kb"
	MB unitTag = "mb"
	GB unitTag = "gb"
)

// ParseSize parsing input expected to return
// the value, and the unit can be: MB, or KB
func ParseSize(input string) (int, StdUnit, error) {
	inputLen := len(input)
	if inputLen <= 2 {
		return 0, 0, custom.InvalidSizeErr
	}
	// getting the postfix (unit) e.g. MB
	unit := strings.ToLower(input[inputLen-2:])
	value, err := strconv.Atoi(input[:inputLen-2])
	if err != nil {
		return 0, 0, custom.InvalidSizeErr
	}

	switch unitTag(unit) {
	case KB:
		return value, KiloByte, nil
	case MB:
		return value, MegaByte, nil
	case GB:
		return value, GigaByte, nil
	}

	return 0, -1, nil
}
