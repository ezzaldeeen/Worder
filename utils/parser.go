package utils

import (
	"strconv"
	"strings"
	"worder/pkg"
)

type unitTag string

const (
	KB unitTag = "kb"
	MB unitTag = "mb"
	GB unitTag = "gb"
)

type GeneratedFileInfo struct {
	size uint
	unit unitTag
}

// ParseSize parsing input expected to return
// the value, and the unit can be: MB, or KB
// todo: StdUnit -> SizeUnit
// todo: return GeneratedFileInfo instead of (int, StdUnit)
// todo: Parser is responsible to convert to bytes
// todo: should be internal (private)
// todo: don't expose too much information
func ParseSize(input string) (int, StdUnit, error) {
	inputLen := len(input)
	if inputLen <= 2 {
		return 0, 0, pkg.InvalidSizeErr
	}
	// getting the postfix (unit) e.g. MB
	unit := strings.ToLower(input[inputLen-2:])
	value, err := strconv.Atoi(input[:inputLen-2])
	if err != nil {
		return 0, 0, pkg.InvalidSizeErr
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

// todo: define the function where you expect to use it
//func ConvertToByte(value int, unit StdUnit) (int, error) {
//	switch unit {
//	case KiloByte:
//		return value * kb, nil
//	case MegaByte:
//		return value * mb, nil
//	case GigaByte:
//		return value * gb, nil
//	default:
//		return 0, custom.InvalidUnitErr
//	}
//}
