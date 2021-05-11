package enderecos

import "strings"

//TipoDeEndereço verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	logradouros := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoTemUmTipoValido := false

	for _, logradouro := range logradouros {
		if logradouro == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tem um tipo Inválido"
}
