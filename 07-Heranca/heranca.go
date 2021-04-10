package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa    // --> essa é a forma de fazer a herança
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	pessoa1 := pessoa{"Joao", "Silva", 43, 170}
	fmt.Println(pessoa1)
	estudante1 := estudante{pessoa1, "Engenharia", "UFOP"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
	fmt.Println(estudante1.sobrenome)
	fmt.Println(estudante1.idade)
	fmt.Println(estudante1.altura)

}
