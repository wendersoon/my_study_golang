package main

import "fmt"

func main() {

	var v1 int = 100 //declaração normal de variável

	var v2 *int //declaração de um ponteiro para um tipo int

	v2 = &v1 //Atri

	var v3 **int

	v3 = &v2

	var v4 ***int

	v4 = &v3

	fmt.Println(***v4, **v3, *v2, v1)
	fmt.Println(&v4, &v3, &v2, v1)

}
