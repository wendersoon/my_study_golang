package main

import (
	"fmt"
	endereco "introducao/enderecos"
)

func main() {
	tipoEndereco := endereco.TipoDeEndereco("Avenida Newton Belo")
	fmt.Println(tipoEndereco)
}
