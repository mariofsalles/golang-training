package main

import (
	"fmt"
	enderecos "introducao-testes/19.1.1-enderecos"
)

func main()  {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

}

// o comando go test irá executar os testes que estão no pacote main
// para executar todos os testes deve-se utilizar o comando: 
// go test ./...
// testes em paralelo: t.paralell
// gerar relatório:
// go test --coverprofile cobertura.txt
// para ler: 
// go tool cover --func=cobertura.txt
// saber o que não está coberto, existe uma função que destaca num relatório
// go tool cover --html=cobertura.txt