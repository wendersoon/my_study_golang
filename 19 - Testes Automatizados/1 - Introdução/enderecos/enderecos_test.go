package endereco

import (
	"testing"
)

type cenariaDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarios := []cenariaDeTeste{
		{"avenida 1", "Avenida"},
		{"Rua da zebra", "Rua"},
		{"Estrada do arroz", "Estrada"},
		{"rodovia perimetral", "Rodovia"},
		{"RUA azul", "Rua"},
		{"AAAAA AAA", "Tipo Inválido"},
	}

	for _, cenario := range cenarios {

		tipoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoRecebido != cenario.tipoEsperado {
			t.Errorf("O tipo recebeido %s é diferente do tipo esperado %s",
				tipoRecebido,
				cenario.tipoEsperado)
		}

	}

}
