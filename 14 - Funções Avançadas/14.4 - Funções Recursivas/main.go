package main

import "fmt"

func fibonnaci(posicao float64) float64 {

	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

func main() {

	posicao := float64(100)

	for i := float64(1); i <= posicao; i++ {
		fmt.Println(fibonnaci(i))
	}
}
