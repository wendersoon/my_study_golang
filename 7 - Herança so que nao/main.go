package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Wenderson", "Melo", 25, 1.99}
	fmt.Println(p1)

	e1 := estudante{p1, "BCC", "IFMA"}
	fmt.Println(e1)
}
