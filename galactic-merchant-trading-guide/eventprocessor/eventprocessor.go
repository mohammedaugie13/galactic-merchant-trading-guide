package eventprocessor

import (
	"fmt"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/currency"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/parser"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/translator"
	"strconv"
)

type EventProcessor struct {
	parser     *parser.Parser
	translator *translator.Translator
	currencies *currency.Currencies
}

func (e *EventProcessor) ProcessStatement(statement string) (string, error) {
	statementType := e.parser.CheckTypeStatement(statement)
	switch statementType {
	case parser.ForeignToRoman:
		foreignSymbol, romanSymbol := e.parser.GetForeignToRoman(statement)
		err := e.translator.AddWord(foreignSymbol, romanSymbol)
		if err != nil {
			return "", err
		}
		return "", nil
	case parser.Credits:
		var err error
		foreignSymbol, creditType, credits := e.parser.GetStatemenCredit(statement)
		value, err := e.translator.ToArabic(foreignSymbol)
		if err != nil {
			return "", err
		}
		creditsInt, _ := strconv.Atoi(credits)
		creditValue := float64(creditsInt) / float64(value)
		err = e.currencies.AddCurrency(creditType, creditValue)
		if err != nil {
			return "", err
		}
		return "", nil
	case parser.HowMuch:
		var err error
		foreignSymbol := e.parser.GetStatementHowMuch(statement)
		value, err := e.translator.ToArabic(foreignSymbol)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%v is %v", foreignSymbol, value), nil
	case parser.HowMany:
		var err error
		foreignSymbol, credits := e.parser.GetStatementHowMany(statement)
		value, err := e.translator.ToArabic(foreignSymbol)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		creditValue, err := e.currencies.GetCurrency(credits)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		totalValue := float64(value) * creditValue
		return fmt.Sprintf("%v %v is %v Credits", foreignSymbol, credits, totalValue), nil
	case parser.HasMore:
		var err error
		foreignSymbolX, creditsX, foreignSymbolY, creditsY := e.parser.GetStatementHasMore(statement)
		valueX, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		creditValueX, err := e.currencies.GetCurrency(creditsX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		totalValueX := float64(valueX) * creditValueX

		valueY, err := e.translator.ToArabic(foreignSymbolY)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		creditValueY, err := e.currencies.GetCurrency(creditsY)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		totalValueY := float64(valueY) * creditValueY

		if totalValueX > totalValueY {
			return fmt.Sprintf("%v %v has more Credits than %v %v", foreignSymbolX, creditsX, foreignSymbolY, creditsY), nil
		}
		return fmt.Sprintf("%v %v has less Credits than %v %v", foreignSymbolX, creditsX, foreignSymbolY, creditsY), nil
	case parser.HasLess:
		var err error
		foreignSymbolX, creditsX, foreignSymbolY, creditsY := e.parser.GetStatementHasLess(statement)
		valueX, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		creditValueX, err := e.currencies.GetCurrency(creditsX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		totalValueX := float64(valueX) * creditValueX

		valueY, err := e.translator.ToArabic(foreignSymbolY)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		creditValueY, err := e.currencies.GetCurrency(creditsY)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		totalValueY := float64(valueY) * creditValueY

		if totalValueX < totalValueY {
			return fmt.Sprintf("%v %v has less Credits than %v %v", foreignSymbolX, creditsX, foreignSymbolY, creditsY), nil
		}
		return fmt.Sprintf("%v %v has more Credits than %v %v", foreignSymbolX, creditsX, foreignSymbolY, creditsY), nil
	case parser.LargerThan:
		var err error
		foreignSymbolX, foreignSymbolY := e.parser.GetStatementLargerThan(statement)
		valueX, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		valueY, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}

		if valueX > valueY {
			return fmt.Sprintf("%v is larger than %v", foreignSymbolX, foreignSymbolY), nil
		}
		return fmt.Sprintf("%v is smaller than %v", foreignSymbolX, foreignSymbolY), nil
	case parser.SmallerThan:
		var err error
		foreignSymbolX, foreignSymbolY := e.parser.GetStatementSmallerThan(statement)
		valueX, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}
		valueY, err := e.translator.ToArabic(foreignSymbolX)
		if err != nil {
			return "", fmt.Errorf("Requested number is in invalid format")
		}

		if valueX < valueY {
			return fmt.Sprintf("%v is smaller than %v", foreignSymbolX, foreignSymbolY), nil
		}
		return fmt.Sprintf("%v is larger than %v", foreignSymbolX, foreignSymbolY), nil

	case parser.NotDefined:
		return "I have no idea what you are talking about", nil
	default:
		return "I have no idea what you are talking about", nil

	}

	return "", nil
}

func NewEventProcessor(t *translator.Translator, p *parser.Parser, c *currency.Currencies) *EventProcessor {
	return &EventProcessor{translator: t, currencies: c, parser: p}
}
