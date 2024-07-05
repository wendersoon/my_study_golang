package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Olá Mundo"))
}

func teste(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Teste"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/", teste)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
