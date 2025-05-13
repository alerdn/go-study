package main

import (
	"os"
	"text/template"
)

type Product struct {
	Description string
	Price       float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	product := Product{
		Description: "banana",
		Price:       1.99,
	}

	file, _ := os.Create("index.html")
	defer file.Close()

	tpl.Execute(file, product)
}
