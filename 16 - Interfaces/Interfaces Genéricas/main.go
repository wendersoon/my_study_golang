package main

import "fmt"

func print(inter ...interface{}) {
	fmt.Println(inter...)
}

func main() {
	print("Olá mundo!")
}
