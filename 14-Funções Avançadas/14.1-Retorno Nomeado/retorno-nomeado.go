package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("Funções com retorno nomeado")
	soma, subtracao := calculosMatematicos(10,20)
	fmt.Printf("Soma: %v e subtração: %v ",soma,subtracao)
}
