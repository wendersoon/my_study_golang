package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loops")

	i := 0

	for i < 10 {
		fmt.Println(i)
		i += 2
	}

	fmt.Println("------------------------")

	for j := 0; j <= 10; j += 3 {
		fmt.Println(j)
	}

	fmt.Println("------------------------")

	slice := []int{1, 2, 3}

	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}

	fmt.Println("------------------------")
	map1 := map[int]string{
		1: "Nome",
		2: "Sobrenome",
	}

	for chave, valor := range map1 {
		fmt.Println(chave, valor)
	}

	fmt.Println("------------------------")

	nome := "Abacate"
	for indice, letra := range nome {
		fmt.Println(indice, string(letra))
	}

	var valor float64
	valor = 2
	for {

		fmt.Println(valor)
		valor *= 2
	}
}
