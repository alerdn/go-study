package cart

import (
	"fmt"
	"log"
	"shopping-cli/payment"
	"shopping-cli/product"
	"shopping-cli/util"
)

var pixStrategy = payment.PixPaymentStrategy{
	Wallet: 100.,
}

var cardStrategy = payment.CardPaymentStrategy{
	Limit: 500.,
}

type cartService struct {
	cart              Cart
	productRepository product.ProductRepository
}

func (cs *cartService) AddProduct() {
	fmt.Println("\n" + util.CenterPad(" Adicionar Produto ", 40, "="))
	fmt.Print("ID do produto: ")

	var id int
	fmt.Scanln(&id)
	id--

	product, err := cs.productRepository.FindByID(id)
	if err != nil {
		log.Println(err)
		return
	}

	cs.cart.AddProduct(product)
	fmt.Println(util.CenterPad(" Carrinho ", 40, "="))
	cs.cart.ShowProducts()
}

func (cs *cartService) RemoveProduct() {
	fmt.Println("\n" + util.CenterPad(" Remover Produto ", 40, "="))
	fmt.Print("ID do produto: ")

	var id int
	fmt.Scanln(&id)
	id--

	cs.cart.RemoveProduct(id)
	fmt.Println(util.CenterPad(" Carrinho ", 40, "="))
	cs.cart.ShowProducts()
}

func (cs *cartService) Payment() {
	fmt.Println("\n" + util.CenterPad(" Pagamento ", 40, "="))
	fmt.Println(util.CenterPad(" Carrinho ", 40, "="))
	cs.cart.ShowProducts()
	fmt.Println("1 - Pix\n2 - Cartão")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		err := cs.cart.Pay(&pixStrategy)
		if err != nil {
			log.Println(err)
			return
		}
	case 2:
		err := cs.cart.Pay(&cardStrategy)
		if err != nil {
			log.Println(err)
			return
		}
	default:
		log.Println("método de pagamento inválido")
	}

	cs.cart.productList = []product.Product{}
}

// Exportado
var CartService = cartService{
	productRepository: product.ProductRepository{},
}
