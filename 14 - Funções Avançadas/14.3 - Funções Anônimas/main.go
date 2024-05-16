package main

import "fmt"

func main() {

	//Retornar nome + s
	nomeMaisS := func(nome string) string {
		return nome + "s"
	}("Abacate")

	fmt.Println(nomeMaisS)

	func(numeros ...int) {
		valor := 1
		for _, i := range numeros {
			valor *= i
		}
		fmt.Println(valor)
	}(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13)

}
