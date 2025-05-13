package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greeting() string {
	return fmt.Sprintf("My name is %s and i'm %d", p.Name, p.Age)
}

func (p Person) Log(s string) string {
	return fmt.Sprintf("%s - %s", time.Now().Format("02/01/2006 03:04"), s)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	p := Person{
		"Alexandre",
		25,
	}

	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p)
}
