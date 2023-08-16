package eventprocessor

import (
	"fmt"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/currency"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/parser"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/roman"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/translator"
	"testing"
)

func TestEventProcessor(t *testing.T) {
	p := parser.NewParser()
	c := currency.NewCurrencies()
	r := roman.NewRoman()
	tr := translator.NewTranslator(r)

	ep := NewEventProcessor(tr, p, c)
	type statement struct {
		stmt string
		res  string
		err  error
	}

	statements := []statement{
		statement{"glob is I", "", nil},
		statement{"prok is V", "", nil},
		statement{"pish is X", "", nil},
		statement{"tegj is L", "", nil},
		statement{"glob glob Silver is 34 Credits", "", nil},
		statement{"glob prok Gold is 57800 Credits", "", nil},
		statement{"pish pish Iron is 3910 Credits", "", nil},
		statement{"how much is pish tegj glob glob ?", "pish tegj glob glob is 42", nil},
		statement{"how many Credits is glob prok Silver ?", "glob prok Silver is 68 Credits", nil},
		statement{"how many Credits is glob glob Gold ?", "glob glob Gold is 28900 Credits", nil},
		statement{"how many Credits is glob glob glob glob glob glob Gold ?", "", fmt.Errorf("Requested number is in invalid format")},
		statement{"how many Credits is pish tegj glob Iron ?", "pish tegj glob Iron is 8015.5 Credits", nil},
		statement{"Does pish tegj glob glob Iron has more Credits than glob glob Gold ?", "pish tegj glob glob Iron has less Credits than glob glob Gold", nil},
		statement{"Does glob glob Gold has less Credits than pish tegj glob glob Iron?", "glob glob Gold has more Credits than pish tegj glob glob Iron", nil},
		statement{"Is glob prok larger than pish pish?", "glob prok is smaller than pish pish", nil},
		statement{"Is tegj glob glob smaller than glob prok?", "tegj glob glob is larger than glob prok", nil},
		statement{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", "I have no idea what you are talking about", nil},
	}

	for _, s := range statements {
		res, _ := ep.ProcessStatement(s.stmt)
		if s.res != res {
			t.Errorf("TestEventProcessor failed, expected: '%v', got: '%v'", s.res, res)
		}

	}
}
