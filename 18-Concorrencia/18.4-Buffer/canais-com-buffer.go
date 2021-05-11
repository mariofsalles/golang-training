package main

import "fmt"

func main() {
	canal := make(chan string, 3) // o 3 é o buffer

	canal <- "olá 1" //envia o valor
	canal <- "olá 2" //envia o valor
	canal <- "olá 3" //envia o valor

	mensagem1 := <-canal // recebe o valor
	mensagem2 := <-canal // recebe o valor
	mensagem3 := <-canal // recebe o valor

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)

}
