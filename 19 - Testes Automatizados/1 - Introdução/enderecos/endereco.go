package endereco

import "strings"

// TipodeEndereco verifica se um endereço é válido
func TipoDeEndereco(endereco string) string {
	tipoValidos := []string{"rua", "avenida", "rstrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	palavra := strings.ToLower(primeiraPalavraDoEndereco)

	enderecoTemUmTipoValido := false

	for _, tipo := range tipoValidos {
		if tipo == palavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return palavra
	}

	return "Tipo Inválido"
}
