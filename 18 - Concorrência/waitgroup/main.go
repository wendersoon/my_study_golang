package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mensagens sync.WaitGroup

	mensagens.Add(2) //Aqui diz quantas goroutines esse grupo terá

	go func() {
		escrever("Olá Mundo!")
		mensagens.Done() // Aqui diz ao waitgroup que a uma das goroutines encerram-se
	}()

	go func() {
		escrever("goroutines é massa demais!")
		mensagens.Done()
	}()

	mensagens.Wait() // Aqui diz para esperar o termino de execução das goroutines
	//O Wait() serve como um ponto de espera no programa

	mensagens.Add(2) //Aqui crei um novo WaitGroup que só será executado ao termino da primeira declarada acima
	go func() {
		escrever("Teste")
		mensagens.Done() // Aqui diz ao waitgroup que a uma das goroutines encerram-se
	}()

	go func() {
		escrever("Teste 2")
		mensagens.Done()
	}()

	mensagens.Wait()
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
