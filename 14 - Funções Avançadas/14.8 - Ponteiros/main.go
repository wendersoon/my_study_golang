package main

import "fmt"

func somaMaisUm(n1 *int) {

	*n1 = *n1 + 1 //Processo de desreferênciação, permite usar o valor do ponteiro

}

func main() {
	valor := 100
	somaMaisUm(&valor) //Passa o endereço de memória para a função
	somaMaisUm(&valor)
	somaMaisUm(&valor)
	somaMaisUm(&valor)
	somaMaisUm(&valor)
	somaMaisUm(&valor)

	fmt.Println(valor) //O valor da variável vai mudar toda vez que a função chamada como ela

}
