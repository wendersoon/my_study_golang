package main

import "fmt"

func Calcular(n1, n2 int) (divisao float64, multiplicacao, soma, subtracao int) {

	divisao = float64(n1) / float64(n2)
	multiplicacao = n1 * n2
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {
	fmt.Println(Calcular(10, 23))
}
