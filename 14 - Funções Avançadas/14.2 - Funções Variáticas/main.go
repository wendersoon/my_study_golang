package main

import (
	"fmt"
)

// Função variática
func multiplicar(numeros ...int) (multiplicacao int) {

	total := 1
	for _, i := range numeros {
		total *= i
	}
	return total
}

func main() {
	resultado := multiplicar(2, 4, 6, 8, 10, 12, 14, 15, 16)
	fmt.Println(resultado)
}
