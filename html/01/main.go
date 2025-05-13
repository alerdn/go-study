package main

import (
	"log"
	"os"
	"text/template"
)

// Tpl é um container que guarda todos os textos/html parseados
var tpl *template.Template

// Roda antes da main
func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}
