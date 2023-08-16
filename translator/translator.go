package translator

import (
	"fmt"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/roman"
	"strings"
)

// Translator for foreign galactic language to human normal language
type Translator struct {
	RomanTranslator *roman.Roman
	Dictionary      map[string]string
}

// AddWord from foreign word  to dictionary
func (t *Translator) AddWord(foreignWord, romanSymbol string) error {
	_, err := t.GetWord(foreignWord)
	if err == nil {
		return fmt.Errorf("%s word already exist", foreignWord)
	}

	t.Dictionary[foreignWord] = romanSymbol

	return nil
}

// ToArabic convert foreign word to arabic numeral
func (t *Translator) ToArabic(foreignWords string) (int, error) {
	words := strings.Split(foreignWords, " ")
	var romans []string

	for _, word := range words {
		roman, err := t.GetWord(word)
		if err != nil {
			return 0, err
		}

		romans = append(romans, roman)
	}

	symbol := strings.Join(romans, "")

	return t.RomanTranslator.ToArabic(symbol)
}

func (t *Translator) GetWord(word string) (string, error) {
	translatedWord, ok := t.Dictionary[word]
	if !ok {
		return "", fmt.Errorf("Word %s not found in our dictionary", word)
	}

	return translatedWord, nil
}

func NewTranslator(roman *roman.Roman) *Translator {
	dictionary := make(map[string]string)

	return &Translator{roman, dictionary}
}
