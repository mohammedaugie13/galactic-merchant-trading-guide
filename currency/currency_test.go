package currency

import (
	"errors"
	"testing"
)

func TestAddNewCurrency(t *testing.T) {
	type currencyTypeTests struct {
		name   string
		amount float64
		err    error
	}

	tests := []currencyTypeTests{
		currencyTypeTests{"Silver", 17, nil},
		currencyTypeTests{"Gold", 14450, nil},
		currencyTypeTests{"Iron", 195.5, nil},
		currencyTypeTests{"Silver", 15, errors.New("Silver currency already exist")},
		currencyTypeTests{"Iron", 200, errors.New("Iron currency already exist")},
	}

	currency := NewCurrencies()

	for _, tt := range tests {
		err := currency.AddCurrency(tt.name, tt.amount)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddCurrency failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestAddSomeCurrencyAndGetCurrency(t *testing.T) {
	type currencyType struct {
		name   string
		amount float64
	}

	type currencyTypeTests struct {
		name string
		err  error
	}

	currencyTypes := []currencyType{
		currencyType{"Silver", 17},
		currencyType{"Gold", 14450},
		currencyType{"Iron", 195.5},
	}

	tests := []currencyTypeTests{
		currencyTypeTests{"Silver", nil},
		currencyTypeTests{"Gold", nil},
		currencyTypeTests{"Iron", nil},
		currencyTypeTests{"Great", errors.New("Currency Great not found in our database")},
		currencyTypeTests{"Good", errors.New("Currency Good not found in our database")},
	}

	currency := NewCurrencies()

	for _, ct := range currencyTypes {
		currency.AddCurrency(ct.name, ct.amount)
	}

	for _, tt := range tests {
		_, err := currency.GetCurrency(tt.name)

		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf("TestAddSomeCurrencyAndGetCurrency failed, expected: '%s', got: '%s'", tt.err.Error(), err.Error())
		}
	}
}

func TestCalculateCurrency(t *testing.T) {
	type currencyType struct {
		name   string
		amount float64
	}

	type currencyTypeTests struct {
		name string
		num  int
		res  float64
	}

	currencyTypes := []currencyType{
		currencyType{"Silver", 17},
		currencyType{"Gold", 14450},
		currencyType{"Iron", 195.5},
	}

	tests := []currencyTypeTests{
		currencyTypeTests{"Silver", 2, 34},
		currencyTypeTests{"Gold", 4, 57800},
		currencyTypeTests{"Iron", 20, 3910},
	}

	currency := NewCurrencies()

	for _, ct := range currencyTypes {
		currency.AddCurrency(ct.name, ct.amount)
	}

	for _, tt := range tests {
		credit, _ := currency.GetCurrency(tt.name)
		res := float64(tt.num) * credit

		if res != tt.res {
			t.Errorf("TestAddSomeCurrencyAndGetCurrency failed, expected: '%f', got: '%f'", tt.res, res)
		}
	}
}
