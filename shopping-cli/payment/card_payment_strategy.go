package payment

import (
	"fmt"
	"shopping-cli/product"
	"shopping-cli/util"
)

type CardPaymentStrategy struct {
	Limit float64
}

func (c *CardPaymentStrategy) Pay(productList []product.Product) error {
	total := .0
	for _, product := range productList {
		total += product.Price
	}

	if c.Limit < total {
		return fmt.Errorf("limite insuficiente")
	}
	c.Limit -= total

	fmt.Printf("\nPagamento efetuado no valor de R$ %05.2f\n%s\n", total, util.BetweenPad("Limite Atual", fmt.Sprintf("R$ %05.2f", c.Limit), 40, "."))
	return nil
}
