package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Olá Mundo!")
	escrever("goroutines é massa demais!")
}

func escrever(texto string) {

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
