package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		//strings.Title = deprecated
		//primeiraPalavraDoEnderecoTitle := strings.Title(primeiraPalavraDoEndereco)
		caser := cases.Title(language.BrazilianPortuguese)
		primeiraPalavraDoEnderecoTitle := caser.String(primeiraPalavraDoEndereco)

		return primeiraPalavraDoEnderecoTitle
	}

	return "Tipo inválido!"
}
