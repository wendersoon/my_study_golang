package main

import "fmt"

func maiorNumero(numeros ...int) int {
	defer fmt.Println("O maior número é: ") //é executado imediatamente antes do return

	valor := numeros[0] //Receber o primeiro número
	for _, i := range numeros {
		if i >= valor {
			valor = i
		}
	}

	return valor
}

func main() {
	//A função não tem nenhum retorno, mas o bloco de código
	// na clausula defer será executado imediatamente antes de terminar a função
	defer fmt.Println("Código encerrado!")

	numeros := []int{1, 2, 12, 124, 12315, 9990, 7, 4}

	resultado := maiorNumero(numeros...)

	fmt.Println(resultado)

}
