package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
