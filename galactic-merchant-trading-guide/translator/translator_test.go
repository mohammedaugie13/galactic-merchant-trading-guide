package translator

import (
	"errors"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/roman"
	"testing"
)

func TestAddWords(t *testing.T) {
	type wordsUnitTest struct {
		words string
		symb  string
		err   error
	}

	tests := []wordsUnitTest{
		wordsUnitTest{"glob", "I", nil},
		wordsUnitTest{"prok", "V", nil},
		wordsUnitTest{"pish", "X", nil},
		wordsUnitTest{"tegj", "X", nil},
		wordsUnitTest{"glob", "I", errors.New("glob word already exist")},
		wordsUnitTest{"pish", "X", errors.New("pish word already exist")},
	}

	translator := NewTranslator(roman.NewRoman())

	for _, tt := range tests {
		err := translator.AddWord(tt.words, tt.symb)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddWords failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestAddSomeWordsAndGetTheirArabicValue(t *testing.T) {
	translator := NewTranslator(roman.NewRoman())
	translator.AddWord("glob", "I")
	translator.AddWord("prok", "V")
	translator.AddWord("pish", "X")
	translator.AddWord("tegj", "L")

	type testWordArabic struct {
		in  string
		out int
		err error
	}

	tests := []testWordArabic{
		testWordArabic{"pish pish prok", 25, nil},
		testWordArabic{"pish tegj glob glob", 42, nil},
	}

	for _, tt := range tests {
		res, _ := translator.ToArabic(tt.in)

		if res != tt.out {
			t.Errorf("TestAddSomeWordsAndGetTheirArabicValue failed, expected: '%d', got: '%d'", tt.out, res)
		}
	}
}

func TestAddSomeWordAndGetErrorWhereWordNotExist(t *testing.T) {
	translator := NewTranslator(roman.NewRoman())
	translator.AddWord("glob", "I")
	translator.AddWord("prok", "V")
	translator.AddWord("pish", "X")
	translator.AddWord("tegj", "L")

	type testWordArabic struct {
		in  string
		out int
		err error
	}

	tests := []testWordArabic{
		testWordArabic{"good great prok", 25, errors.New("Word good not found in our dictionary")},
		testWordArabic{"pish great nice glob", 42, errors.New("Word great not found in our dictionary")},
	}

	for _, tt := range tests {
		_, err := translator.ToArabic(tt.in)

		if err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeWordAndGetErrorWhereWordNotExist failed, expected: '%s', got: '%s'",
				tt.err.Error(), err.Error())
		}
	}
}
