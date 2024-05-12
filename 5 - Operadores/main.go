package main

import "fmt"

func main() {
	//Aritméticos
	soma := 5 + 5
	subtracao := 5 - 0
	divisao := 5 / 5
	multiplicacao := 5 * 5
	modulo := 5 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, modulo)

	// Atribuição
	var nome string = "String"
	nome2 := "String 2"

	fmt.Println(nome, nome2)

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores Lógicos
	fmt.Println("-----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //and
	fmt.Println(verdadeiro || falso) //or
	fmt.Println(!verdadeiro)         //not

	//Operadores Unários
	numero := 10
	numero++
	fmt.Println(numero)

	numero += 100
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero -= 100
	fmt.Println(numero)

	numero *= 10
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 10
	fmt.Println(numero)

	//Operadores Ternários
	//Não tem

}
