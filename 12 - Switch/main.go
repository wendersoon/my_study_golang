package main

import "fmt"

func verSeNumeroPar(numero int) bool {
	switch {
	case numero%2 == 0:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println("Ver se numero é par ou impar")

	numero := 1023423412542365

	result := verSeNumeroPar(numero)

	if result {
		fmt.Println("Número é par.")
	} else {
		fmt.Println("Número é ímpar.")
	}

}
