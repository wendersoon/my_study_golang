package main

import "fmt"

func contador(numero int) func() int {

	resultado := func() int {
		numero++
		return numero
	}
	return resultado
}

func main() {
	contar := contador(10)

	fmt.Println(contar()) // a variavel interna da função contador não é descartada quando a função
	fmt.Println(contar()) //termina a execução mas seu estado sem mantêm

	contar2 := contador(1) //Os clusures são independentes entre si
	fmt.Println(contar2())
	fmt.Println(contar2())
	fmt.Println(contar2())

}
