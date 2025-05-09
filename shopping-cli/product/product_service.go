package product

import (
	"fmt"
	"shopping-cli/util"
)

type productService struct {
	repository ProductRepository
}

func (ps *productService) ShowProducts() {
	fmt.Println("\n" + util.CenterPad(" Lista de Produtos ", 40, "="))
	for i, product := range ps.repository.FindAll() {
		fmt.Printf("%d - %s\n", i+1, util.BetweenPad(product.Description, fmt.Sprintf("R$ %05.2f", product.Price), 36, " "))
	}
}

func (ps *productService) FindByID(id int) (Product, error) {
	return ps.repository.FindByID(id)
}

// Exportado
var ProductService = productService{
	repository: ProductRepository{},
}
