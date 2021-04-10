package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("***Arquivo Structs")

	var u usuario
	u.nome = "Mario"
	u.idade = 41
	fmt.Println(u)

	var end endereco
	end.logradouro = "Rua Teste"
	end.numero = 125
	// end := endereco {"Rua Teste", 125}

	u2 := usuario{"Mario", 41, end}
	fmt.Println(u2)

	u3 := usuario{nome: "Mario"}
	fmt.Println(u3)

	u3 = usuario{idade: 41}
	fmt.Println(u3)
}
