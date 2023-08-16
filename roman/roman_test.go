package roman

import (
	"errors"
	"testing"
)

type romanTests struct {
	in  string
	out int
	err error
}

func TestToArabic(t *testing.T) {
	tests := []romanTests{
		romanTests{"XXXIX", 39, nil},
		romanTests{"MVI", 1006, nil},
		romanTests{"MMVI", 2006, nil},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToArabic(tt.in)

		if err != tt.err {
			t.Error(err)
		} else if num != tt.out {
			t.Errorf("TestToNumber failed, expected: '%d', got: '%d'", tt.out, num)
		}
	}
}

func TestSmallValueSymbolSubtracted(t *testing.T) {
	tests := []romanTests{
		romanTests{"XLV", 45, nil},
		romanTests{"LXXXIX", 89, nil},
		romanTests{"MCMXLIV", 1944, nil},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToArabic(tt.in)

		if err != tt.err {
			t.Error(err)
		} else if num != tt.out {
			t.Errorf("TestSmallValueSymbolSubtracted failed, expected: '%d', got: '%d'", tt.out, num)
		}
	}
}

func TestCanBeRepeatedUntilThreeTimesForParticularSymbols(t *testing.T) {
	tests := []romanTests{
		romanTests{"III", 3, nil},
		romanTests{"XXXIX", 39, nil},
		romanTests{"CCCC", 0, errors.New("This roman number C can't be repeated more than 3 times")},
		romanTests{"CXXXX", 0, errors.New("This roman number X can't be repeated more than 3 times")},
		romanTests{"MMMM", 0, errors.New("This roman number M can't be repeated more than 3 times")},
	}

	roman := NewRoman()

	for _, tt := range tests {

		num, err := roman.ToArabic(tt.in)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestCanBeSubtracted failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		} else if num != tt.out {
			t.Errorf("TestCanBeSubtracted failed, expected: '%d', got: '%d'", tt.out, num)
		}
	}
}

func TestCanBeSubtractedForParticularSymbols(t *testing.T) {
	tests := []romanTests{
		romanTests{"IV", 4, nil},
		romanTests{"IX", 9, nil},
		romanTests{"XLV", 45, nil},
		romanTests{"XCIV", 94, nil},
		romanTests{"CDX", 410, nil},
		romanTests{"CMXL", 940, nil},
		romanTests{"IL", 0, errors.New("I roman number can't be substracted with L")},
		romanTests{"XXD", 0, errors.New("X roman number can't be substracted with D")},
		romanTests{"XMLX", 0, errors.New("X roman number can't be substracted with M")},
	}

	roman := NewRoman()

	for _, tt := range tests {
		num, err := roman.ToArabic(tt.in)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestCanBeSubtractedForParticularSymbols failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		} else if num != tt.out {
			t.Errorf("TestCanBeSubtractedForParticularSymbols failed, expected: '%d', got: '%d'", tt.out, num)
		}
	}
}
