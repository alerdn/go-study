package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Phone int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	file, _ := os.Create("index.html")
	defer file.Close()

	contacts := []Person{
		{"John Doe", 123456789},
		{"Jane Smith", 987654321},
		{"Alice Johnson", 456789123},
	}

	tpl.Execute(file, contacts)
}
