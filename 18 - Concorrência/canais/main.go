package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)

	go escrever("Hello!", canal)

	//Uma forma de obter todos os valores de um canal até ele fechar
	// for mensagem := range canal {
	// 	fmt.Println((mensagem))
	// }

	// Outra forma de obter todos os valores e encerrando quando o canal é fechado
	for {
		mensagem, open := <-canal
		if !open {
			break
		}
		fmt.Println(mensagem)
	}

}
func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	//Fechar Canal
	close(canal)
}
