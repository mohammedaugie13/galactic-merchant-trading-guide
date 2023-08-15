package parser

import (
	"testing"
)

func TestStatementType(t *testing.T) {
	type statement struct {
		stmt string
		Type int8
	}

	statements := []statement{
		statement{"glob is I", ForeignToRoman},
		statement{"prok is V", ForeignToRoman},
		statement{"pish is X", ForeignToRoman},
		statement{"tegj is L", ForeignToRoman},
		statement{"glob glob Silver is 34 Credits", Credits},
		statement{"glob prok Gold is 57800 Credits", Credits},
		statement{"pish pish Iron is 3910 Credits", Credits},
		statement{"how much is pish tegj glob glob ?", HowMuch},
		statement{"how many Credits is glob prok Silver ?", HowMany},
		statement{"how many Credits is glob glob Gold ?", HowMany},
		statement{"how many Credits is glob glob glob glob glob glob Gold ?", HowMany},
		statement{"how many Credits is pish tegj glob Iron ?", HowMany},
		statement{"Does pish tegj glob glob Iron has more Credits than glob glob Gold ?", HasMore},
		statement{"Does glob glob Gold has less Credits than pish tegj glob glob Iron?", HasLess},
		statement{"Is glob prok larger than pish pish?", LargerThan},
		statement{"Istegj glob glob smaller than glob prok?", SmallerThan},
		statement{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", NotDefined},
	}

	parser := NewParser()

	for _, s := range statements {
		res := parser.CheckTypeStatement(s.stmt)
		if s.Type != res {
			t.Errorf("TestTypeStatement failed, expected: '%v', got: '%v'", s.Type, res)
		}

	}
}
