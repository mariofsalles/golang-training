package main

import "fmt"

func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

// o parametro variatico obrigatoriamente deve ser o ultimo e só pode ser unico por função
func escrever(texto string, numeros ...int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}


func main() {
	range1 := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(range1)

	range2 := soma()
	fmt.Println(range2)

	escrever("Olá", 10, 2, 3, 4, 5, 7)

}
