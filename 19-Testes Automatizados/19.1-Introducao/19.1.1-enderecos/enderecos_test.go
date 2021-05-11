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
	//TestXXXXX
	enderecoParaTeste := "Rua ABC"

	TipoDeEnderecoEsperado := "Avenida"
	TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if TipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Errorf("\n\nO tipo recebido é diferente do esperado! \nEsperava: " +
			strings.ToUpper(TipoDeEnderecoEsperado) +
			" e recebeu: " +
			strings.ToUpper(TipoDeEnderecoRecebido) + "\n\n")
	}
}
