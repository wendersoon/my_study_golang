package main

import "fmt"

type endereco struct {
	rua    string
	numero int8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type formadePagamento struct {
	dinheiro bool
	cartao   bool
	usuario  usuario
}

func main() {
	fmt.Println("Arquivo Structs")

	//Uma forma de declarar uma struct
	var u usuario
	u.nome = "Wenderson"
	u.idade = 25
	fmt.Println(u)

	//Declarando um struct com inferência de tipo
	u2 := usuario{"Wenderson", 25, endereco{"b", 12}} //Aqui crie a struct aninhada diretamente
	fmt.Println(u2)

	pagamento := formadePagamento{false, true, u2}
	fmt.Println(pagamento)

	//Criando struct com valores vazios
	adress := endereco{numero: 12}

	u3 := usuario{nome: "Wenderson", endereco: adress}

	pagamento2 := formadePagamento{cartao: false, usuario: u3}
	fmt.Println(pagamento2)

}
