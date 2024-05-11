package main

import "fmt"

func main() {
	// Declaração explicita
	var nome string = "abacate"
	fmt.Println(nome)

	//Declaração implicíta
	nome2 := "maçã"
	fmt.Println(nome2)

	//Declaração múltipla explicita
	var (
		sobrenome1 string = "Doce"
		sobrenome2 string = "Amargo"
	)
	fmt.Println(sobrenome1, sobrenome2)

	//Declaração múltipla implicita
	first, second := "1", "2"
	fmt.Println(first, second)

	//Inversao de valores
	first, second = second, first
	fmt.Println(first, second)

	//Declaração múltipla explicita sem atribuir valor
	var (
		sobrenome3 string
		sobrenome4 string
	)
	fmt.Println(sobrenome3, sobrenome4)

	//Atribuindo valores
	sobrenome3 = "Abacaxi"
	sobrenome4 = "Melancia"
	fmt.Println(sobrenome3, sobrenome4)
}
