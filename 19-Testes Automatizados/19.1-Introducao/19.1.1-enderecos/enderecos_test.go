// Teste Unitário
package enderecos

import (
	"strings"
	"testing"
)

type CenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []CenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrante", "Rodovia"},
		{"Praça das Rosas", "Tem um tipo Inválido"},
		{"RUA DOS BOBOS", "Ra"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		tipoDeEnderecoEsperado := cenario.retornoEsperado
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("\n\nO tipo recebido é: " + strings.ToUpper(tipoDeEnderecoRecebido) +
				" esperava: " + strings.ToUpper(tipoDeEnderecoEsperado)+"\n\n")
		}
	}
}
