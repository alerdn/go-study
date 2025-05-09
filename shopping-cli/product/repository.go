package product

import (
	"errors"
)

var productsSource = []Product{
	{Description: "X-Burger", Price: 13.},
	{Description: "X-Calabresa", Price: 20.},
	{Description: "X-Bacon", Price: 20.},
	{Description: "Coca-cola", Price: 9.},
	{Description: "Pizza", Price: 50.},
	{Description: "Suco", Price: 15.},
	{Description: "Pastel de carne", Price: 12.},
	{Description: "Pastel de frango", Price: 15.},
}

type ProductRepository struct {
}

func (pr *ProductRepository) FindAll() []Product {
	return productsSource
}

func (pr ProductRepository) FindByID(id int) (Product, error) {
	if id < 0 || id > len(productsSource) {
		return Product{}, errors.New("produto não encontrado")
	}

	return productsSource[id], nil
}
