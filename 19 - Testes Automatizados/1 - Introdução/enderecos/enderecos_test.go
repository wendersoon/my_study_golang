package endereco

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Newton Belo"

	tipoEsperado := "avenida"
	tipoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoEsperado != tipoRecebido {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoEsperado,
			tipoRecebido)
	}
}
