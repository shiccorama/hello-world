package main

import (
	"html/template"

	"net/http"
)

var tpl *template.Template

type pageData struct {
	YourName string
	PassCode string
	Email    string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/login.html"))
}

func main() {

	http.HandleFunc("/site/", MyServe)
	http.ListenAndServe(":8080", nil)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
}

func MyServe(w http.ResponseWriter, hur *http.Request) {

	tpl.ExecuteTemplate(w, "login.html", nil)
}
