package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	// tpl := template.Must(template.New("greetings").Parse("My name is {{.}}"))
	// tpl.ExecuteTemplate(os.Stdout, "greetings", "alexandre")

	fm := template.FuncMap{
		"upper": strings.ToUpper,
	}

	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	tpl.ExecuteTemplate(os.Stdout,"index.gohtml", "alex")
}
