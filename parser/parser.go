package parser

import (
	"regexp"
)

const (
	ForeignToRoman int8 = 0
	Credits        int8 = 1
	HowMuch        int8 = 2
	HowMany        int8 = 3
	HasMore        int8 = 4
	HasLess        int8 = 5
	LargerThan     int8 = 6
	SmallerThan    int8 = 7
	NotDefined     int8 = 8
)

type Parser struct{}

func (p *Parser) CheckTypeStatement(statement string) int8 {
	isForeignStatement := p.IsStatementForeignToRoman(statement)
	if isForeignStatement {
		return ForeignToRoman
	}
	isCredit := p.IsStatementCredit(statement)
	if isCredit {
		return Credits
	}
	isHowMuch := p.IsStatementHowMuch(statement)
	if isHowMuch {
		return HowMuch
	}
	isHowMany := p.IsStatementHowMany(statement)
	if isHowMany {
		return HowMany
	}
	isHasMore := p.IsStatementHasMore(statement)
	if isHasMore {
		return HasMore
	}
	isHasLess := p.IsStatementHasLess(statement)
	if isHasLess {
		return HasLess
	}
	isLargerThan := p.IsStatementLargerThan(statement)
	if isLargerThan {
		return LargerThan
	}
	isSmallerThan := p.IsStatementSmallerThan(statement)
	if isSmallerThan {
		return SmallerThan
	}

	return NotDefined
}

func (p *Parser) IsStatementForeignToRoman(statement string) bool {
	pattern := `^\w+\s+is\s+\w+$`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)
	if len(match) == 0 {
		return false
	}
	p.GetForeignToRoman(statement)
	return true
}

func (p *Parser) GetForeignToRoman(statement string) (string, string) {
	pattern := `^(\w+)\s+is\s+(\w+)$`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	foreign := match[1]
	roman := match[2]

	return foreign, roman

}

func (p *Parser) IsStatementCredit(statement string) bool {
	pattern := `^(.+) (\w+) is (\d+) Credits$`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)
	if len(match) == 0 {
		return false
	}

	return true
}

func (p *Parser) GetStatemenCredit(statement string) (string, string, string) {
	pattern := `^(.+) (\w+) is (\d+) Credits$`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	foreignSymbol := match[1]
	creditType := match[2]
	credits := match[3]
	return foreignSymbol, creditType, credits
}

func (p *Parser) IsStatementHowMuch(statement string) bool {
	pattern := `how much is`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}

	return true
}

func (p *Parser) GetStatementHowMuch(statement string) string {
	pattern := `how much is (.+)\s+\?`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	return match[1]
}

func (p *Parser) IsStatementHowMany(statement string) bool {
	pattern := `how many Credits is`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}
	p.GetStatementHowMany(statement)

	return true
}

func (p *Parser) GetStatementHowMany(statement string) (string, string) {
	var re *regexp.Regexp
	pattern := `how many Credits is (.+)\s+\?`

	re = regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	pattern2 := `^([\w\s]+) (\w+)`
	re = regexp.MustCompile(pattern2)
	match2 := re.FindStringSubmatch(match[1])
	foreign := match2[1]
	credit := match2[2]
	return foreign, credit
}

func (p *Parser) IsStatementHasMore(statement string) bool {
	pattern := `has more`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}

	return true
}

func (p *Parser) GetStatementHasMore(statement string) (string, string, string, string) {
	var re *regexp.Regexp
	pattern := `Does ([\w\s]+) (\w+) has more Credits than ([\w\s]+) (\w+)`

	re = regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)
	xForeign := match[1]
	xCredit := match[2]
	yForeign := match[3]
	yCredit := match[4]

	return xForeign, xCredit, yForeign, yCredit

}

func (p *Parser) IsStatementHasLess(statement string) bool {
	pattern := `has less`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}
	p.GetStatementHasLess(statement)

	return true
}

func (p *Parser) GetStatementHasLess(statement string) (string, string, string, string) {
	var re *regexp.Regexp
	pattern := `Does ([\w\s]+) (\w+) has less Credits than ([\w\s]+) (\w+)`

	re = regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)
	xForeign := match[1]
	xCredit := match[2]
	yForeign := match[3]
	yCredit := match[4]

	return xForeign, xCredit, yForeign, yCredit

}

func (p *Parser) IsStatementLargerThan(statement string) bool {
	pattern := `larger than`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}

	return true
}

func (p *Parser) GetStatementLargerThan(statement string) (string, string) {
	var re *regexp.Regexp
	pattern := `Is ([\w\s]+) larger than ([\w\s]+)`

	re = regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	return match[1], match[2]

}

func (p *Parser) IsStatementSmallerThan(statement string) bool {
	pattern := `smaller than`

	re := regexp.MustCompile(pattern)
	match := re.FindAllString(statement, -1)
	if len(match) == 0 {
		return false
	}

	return true
}

func (p *Parser) GetStatementSmallerThan(statement string) (string, string) {
	var re *regexp.Regexp
	pattern := `Is ([\w\s]+) smaller than ([\w\s]+)`

	re = regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(statement)

	return match[1], match[2]

}
func NewParser() *Parser {
	return &Parser{}
}
