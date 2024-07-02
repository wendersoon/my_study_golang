package endereco

import "strings"

// TipodeEndereco verifica se um endereço é válido
func TipoDeEndereco(endereco string) string {
	tipoValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	palavra := strings.ToLower(primeiraPalavraDoEndereco)

	enderecoTemUmTipoValido := false

	for _, tipo := range tipoValidos {
		if tipo == palavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(palavra)
	}

	return "Tipo Inválido"
}
