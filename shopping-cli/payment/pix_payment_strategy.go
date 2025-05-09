package payment

import (
	"fmt"
	"shopping-cli/product"
	"shopping-cli/util"
)

type PixPaymentStrategy struct {
	Wallet float64
}

func (p *PixPaymentStrategy) Pay(productList []product.Product) error {
	total := .0
	for _, product := range productList {
		total += product.Price
	}

	if p.Wallet < total {
		return fmt.Errorf("saldo insuficiente")
	}
	p.Wallet -= total

	fmt.Printf("\nGerando QRCode no valor de R$ %05.2f\n%s\n", total, util.BetweenPad("Limite Atual", fmt.Sprintf("R$ %05.2f", p.Wallet), 40, "."))
	return nil
}
