package currency

import "fmt"

type Currencies struct {
	Currency map[string]float64
}

// AddCurrency from new currency
func (c *Currencies) AddCurrency(currency string, amount float64) error {
	_, err := c.GetCurrency(currency)
	if err == nil {
		return fmt.Errorf("%s currency already exist", currency)
	}

	c.Currency[currency] = amount

	return nil
}

// GetCurrency retrieve credit value from specified currency
func (c *Currencies) GetCurrency(currency string) (float64, error) {
	amount, ok := c.Currency[currency]
	if !ok {
		return 0, fmt.Errorf("Currency %s not found in our database", currency)
	}

	return amount, nil
}

func NewCurrencies() *Currencies {
	currencies := make(map[string]float64)

	return &Currencies{Currency: currencies}
}
