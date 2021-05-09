package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	fmt.Println("Funções Recursivas")
	//0 1 1 2 3 5 8 13 21 ...
	posicao := uint(15)
	for i := uint(0); i < posicao; i++{
		fmt.Println(fibonacci(i))
	}
}
