package payment

import "shopping-cli/product"

type IPaymentStrategy interface {
	Pay(productList []product.Product) error
}
