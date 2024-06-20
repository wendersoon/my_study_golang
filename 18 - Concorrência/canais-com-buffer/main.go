package main

import "fmt"

func main() {

	canal := make(chan string, 2)

	canal <- "Oi"
	canal <- "io"

	mensagem := <-canal

	fmt.Println(mensagem)
	mens2 := <-canal
	fmt.Println(mens2)

}
