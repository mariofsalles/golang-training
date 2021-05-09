package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado") // será executado antes do retorno
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// fmt.Println("Sem defer")
	// fmt.Println("---------")
	// funcao1()
	// funcao2()
	// fmt.Println("\nCom defer")
	// fmt.Println("---------")
	// defer funcao1()
	// funcao2()
	fmt.Println("\nCalcula medias")
	fmt.Println("---------")
	_aprovado := alunoAprovado(7,8)
	fmt.Println(_aprovado)
}
