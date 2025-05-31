package main

import (
	"net/http"
	"text/template"
)

var tpl template.Template

type Handler struct{}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
}

func init() {
	tpl = *template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.ListenAndServe(":8080", &Handler{})
}
