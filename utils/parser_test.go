package utils

import (
	"errors"
	"testing"
	"worder/pkg"
)

type expectedParsing struct {
	input    string
	outValue int
	outUnit  StdUnit
}

func TestParseSize(t *testing.T) {
	expectations := []expectedParsing{
		{input: "100MB", outValue: 100, outUnit: MegaByte},
		{input: "80Kb", outValue: 80, outUnit: KiloByte},
		{input: "2gb", outValue: 2, outUnit: GigaByte},
	}

	for _, exp := range expectations {
		actValue, actUnit, err := ParseSize(exp.input)
		if err != nil {
			if !errors.Is(err, pkg.InvalidSizeErr) {
				t.Error(err)
			}
		}
		if actValue != exp.outValue {
			t.Errorf("Expected Value: %d - Actual Value: %d", exp.outValue, actValue)
		}
		if actUnit != exp.outUnit {
			t.Errorf("Expected Unit: %d - Actual Unit: %d", exp.outUnit, actUnit)
		}
	}
}
