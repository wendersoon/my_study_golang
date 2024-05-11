package main

import (
	"errors"
	"fmt"
)

func main() {

	// Inteiros: int8, int16, int32 e int64

	var num1 int8 = 100
	fmt.Println(num1)

	var num2 int16 = 10000
	fmt.Println(num2)

	var num3 int32 = 1000000000
	fmt.Println(num3)

	var num4 int64 = 1000000000000000000
	fmt.Println(num4)

	// Inteiros sem sinal

	var num5 uint8 = 100
	fmt.Println(num5)

	var num6 uint16 = 10000
	fmt.Println(num6)

	var num7 uint32 = 1000000000
	fmt.Println(num7)

	var num8 uint64 = 1000000000000000000
	fmt.Println(num8)

	//Acima todos sã declarações explicitas, abaixo uma declaração implecita
	// o tipo é dado por inferência
	numero := 10
	fmt.Println(numero)

	// Aliás para inteiros

	var num9 rune = 1000000000 //int32
	fmt.Println(num9)

	var num10 byte = 100 //uint8
	fmt.Println(num10)

	//#########################################
	// Tipo float

	var f1 float32 = 1.0
	fmt.Println(f1)

	var f2 float64 = 1.64646
	fmt.Println(f2)

	//########################################
	// Tipo String

	var nome string = "Lua" //Sempre deve ser usados com aspas duplas
	fmt.Println(nome)

	test := 'a' // ao usar aspas simples, o caracter é convertido para o número ASCII
	fmt.Println(test)

	//#######################################
	//Tipo Booleano

	var azul bool = true
	fmt.Println(azul)

	//#######################################
	//Tipo Erro

	var erro error = errors.New("Novo Erro!") //error é um tipo de dado dentro do Go
	//Para ser criado deve usar o pacote nativo errors.New()

	fmt.Println(erro)

	//######################################
	//Valor zero quando a variável não é inicializada

	var sobrenome string //String vazia
	fmt.Println(sobrenome)

	var number int //0
	fmt.Println(number)

	var numberf float64 //0
	fmt.Println(numberf)

	var booleano bool // false
	fmt.Println(booleano)

	var erro2 error //<nil>
	fmt.Println(erro2)

}
