package main

import "fmt"

func verNumero(numero int) {
	if numero > 1000 {
		fmt.Println("Número é maior que 1000.")
	} else if numero >= 50 {
		fmt.Println("Número é maior que 50.")
		if numero <= 75 {
			fmt.Println("Número menor que 75.")
		} else {
			fmt.Println("Número maior que 75.")
			if numero >= 100 {
				fmt.Println("Número maior ou igual 100.")
			}
		}
	} else if numero < 50 {
		fmt.Println("Número é menor que 50.")
	}
}

func main() {
	fmt.Println("Estruturas de Controle")

	//IF, IF...ELSE IF...ELSE
	numero := 100
	verNumero(numero)

	numero2 := 10
	verNumero(numero2)

	//IF...INIT

	if nome := "Tecas"; nome == "Teca" {
		fmt.Println(nome)
	} else {
		fmt.Println("Jabuticaba")
	}

}
