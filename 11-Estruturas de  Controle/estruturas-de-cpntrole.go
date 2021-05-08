package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")
	fmt.Println("----------------------------")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	fmt.Println("----------------------------")

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}

}
