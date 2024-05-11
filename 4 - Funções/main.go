package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func fazer(n1, n2 int, nome, sobrenome string) (string, int) {

	fmt.Println(nome, sobrenome)
	return sobrenome, n1 + n2

}
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Atribuição de função para uma variável
	var nome = func(nome string, sobrenome string) int {
		fmt.Println(nome)
		fmt.Println(sobrenome)
		return 25
	}

	idade := nome("Wenderson", "Melo")
	fmt.Println(idade)

	sobrenome, idade := fazer(15, 25, "Wenderson", "Melo")
	fmt.Println(sobrenome)
	fmt.Println(idade)
}
