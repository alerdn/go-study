package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shopping-cli/cart"
	"shopping-cli/product"
	"shopping-cli/util"
	"strconv"
	"strings"
)

type MenuOption struct {
	Description string
	Action      func()
}

func main() {
	menu := []MenuOption{
		{
			"Listar Produtos",
			func() {
				product.ProductService.ShowProducts()
			},
		},
		{
			"Adicionar Produto",
			func() {
				cart.CartService.AddProduct()
			},
		},
		{
			"Remover Produto",
			func() {
				cart.CartService.RemoveProduct()
			},
		},
		{
			"Ir para pagamento",
			func() {
				cart.CartService.Payment()
			},
		},
		{
			"Sair",
			func() {
				os.Exit(0)
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n" + util.CenterPad(" MENU ", 40, "="))
		for i, menuOption := range menu {
			fmt.Printf("%d - %s\n", i+1, menuOption.Description)
		}

		scanner.Scan()

		optionIndex, err := strconv.Atoi((strings.TrimSpace(scanner.Text())))
		optionIndex--

		if err != nil || (optionIndex < 0 || optionIndex > len(menu)) {
			log.Println("Opção inválida")
			continue
		}

		option := menu[optionIndex]
		option.Action()
	}
}
