package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	map1 := map[int]string{
		1: "Wenderson",
		2: "Flávia",
	}

	fmt.Println(map1)
	//Quero adicionar mais um elemento no map
	map1[3] = "Rafael"
	fmt.Println(map1)

	//Quero editar um nome
	map1[2] = "Flavia"
	fmt.Println(map1)

	//Quero criar um map aninhado

	map2 := map[string]map[int]string{
		"primeiro": {
			1: "Nome",
		},
		"segundo": {
			1: "Sobrenome",
		},
	}
	fmt.Println(map2)

	//Quero remover um elemento
	delete(map2, "primeiro")
	fmt.Println(map2)
}
