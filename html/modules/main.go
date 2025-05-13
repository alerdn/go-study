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
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Alexandre")
}
