package main

import (
	"html/template"

	"net/http"
)

var tpl *template.Template

type pageData struct {
	NameUser string
	Password string
	Email    string
}

func init() {
	tpl = template.Must(template.ParseGlob("NewOne/log.html"))
}

func main() {

	http.HandleFunc("/site/", MyServe)
	http.ListenAndServe(":8080", nil)
	http.Handle("/NewOne/", http.StripPrefix("/NewOne/", http.FileServer(http.Dir("NewOne"))))
}

func MyServe(w http.ResponseWriter, hur *http.Request) {
	pd := pageData{
		NameUser: "",
		Password: "",
	}
	if (hur.Method == http.MethodPost) {
		var NameUser, Password string
		NameUser = hur.FormValue("TxtNameUser")
		Password = hur.FormValue("TxtPasswordUser")
		pd.NameUser = NameUser
		pd.Password = Password
	}
	tpl.ExecuteTemplate(w, "log.html", pd)
}
