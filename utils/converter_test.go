package utils

import (
	"errors"
	"testing"
	"worder/pkg"
)

type expectedConversion struct {
	inValue  int
	inUnit   StdUnit
	outValue int
}

func TestConvertToByte(t *testing.T) {
	wrongUnit := StdUnit(-1)
	expectations := []expectedConversion{
		{inValue: 1000, inUnit: KiloByte, outValue: 1000 * kb},
		{inValue: 100, inUnit: MegaByte, outValue: 100 * mb},
		{inValue: 10, inUnit: GigaByte, outValue: 10 * gb},
		{inValue: 10, inUnit: wrongUnit, outValue: 0},
	}

	for _, exp := range expectations {
		actual, err := ConvertToByte(exp.inValue, exp.inUnit)
		if err != nil {
			if !errors.Is(err, pkg.InvalidUnitErr) {
				t.Errorf("Expected: %v - Actual: %v", exp.outValue, actual)
			}
		}
		if actual != exp.outValue {
			t.Errorf("Expected: %v - Actual: %v", exp.outValue, actual)
		}
	}
}
