package main

import "fmt"

func recupearExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada!")
	}
}

func verSomaNumero(n1, n2 int) bool {
	defer fmt.Println("Iniciando") //Toda clausula defer é executada antes do Panic
	defer recupearExecucao()       //Agora o programa é recuperado do pânico

	resultao := n1 + n2
	if resultao >= 0 {
		return true
	}

	panic("Valor é menor que zero") //O programa entra em pânico, a clausula recover() recupera a execução

}

func main() {

	resultado := verSomaNumero(1, -2)
	fmt.Println(resultado)
}
