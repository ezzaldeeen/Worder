package utils

import "worder/custom"

type StdUnit int

const (
	KiloByte StdUnit = iota
	MegaByte
	GigaByte

	kb = 1024
	mb = kb * 1024
	gb = mb * 1204
)

func ConvertToByte(value int, unit StdUnit) (int, error) {
	switch unit {
	case KiloByte:
		return value * kb, nil
	case MegaByte:
		return value * mb, nil
	case GigaByte:
		return value * gb, nil
	default:
		return 0, custom.InvalidUnitErr
	}
}
