package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	file, _ := os.Create("index.html")
	defer file.Close()

	// products := []string{"banana", "apple", "orange"}
	// tpl.Execute(file, products)

	products := map[string]float64{
		"banana": 1.99,
		"apple":  2.99,
		"orange": 3.99,
	}
	tpl.Execute(file, products)
}
