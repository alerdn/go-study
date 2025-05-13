package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	family := []Person{
		{"Alexandre", 25},
		{"Laynara", 27},
		{"Helena", 1},
	}

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", family)
}
