package cart

import (
	"fmt"
	"shopping-cli/payment"
	"shopping-cli/product"
	"shopping-cli/util"
	"slices"
)

type Cart struct {
	productList []product.Product
}

func (c *Cart) AddProduct(product product.Product) {
	c.productList = append(c.productList, product)
}

func (c *Cart) RemoveProduct(index int) {
	c.productList = slices.Delete(c.productList, index, index+1)
}

func (c *Cart) ShowProducts() {
	total := 0.
	for i, product := range c.productList {
		total += product.Price
		fmt.Printf("%d - %s\n", i+1, util.BetweenPad(product.Description, fmt.Sprintf("R$ %05.2f", product.Price), 36, "."))
	}
	fmt.Println("\n" + util.BetweenPad("Total", fmt.Sprintf("R$ %05.2f", total), 40, "."))
}

func (c *Cart) Pay(ps payment.IPaymentStrategy) error {
	return ps.Pay(c.productList)
}
