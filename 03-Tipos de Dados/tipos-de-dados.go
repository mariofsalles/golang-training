package main

import (
	"errors"
	"fmt"
)

func main(){
	// var numero int64 = 10000000000000000000 ocorre um overflow
	// var numero int32 = 1000000000 ocorre um overflow
	// var numero int16 = 10000 ocorre um overflow
	// var numero int8 = 1000 ocorre um overflow

	fmt.Println("\n-------Inteiros--------")

	var numero1 int8 = 100
	fmt.Println(numero1)
	
	// usa a arquitetura do computador:
	var numero2 int = 1000000000000000000
	fmt.Println(numero2)

	// sem sinal --> uint
	var numero3 uint = 1000000000000000000
	fmt.Println(numero3)

	// alias
	// int32 = rune
	var numero4 rune = 1234565633
	fmt.Println(numero4)
	// int8 = byte
	var numero5 byte = 123
	fmt.Println(numero5)

	fmt.Println("\n-------Reais--------")
	var(
		numero6, numero7 float32 = 1.225,2.33333
		numero8, numero9 float64 = 1.2555555555555,2.55555555555555555555555
	)
	fmt.Print(numero6,"\n",numero7,"\n",numero8,"\n",numero9,"\n ")

	fmt.Println("\n-------Char--------")
	char := 'B'
	fmt.Println(char,"\n ") 

	fmt.Println("\n-------Valor zero--------")
	var texto int
	fmt.Println(texto)

	fmt.Println("\n-------Booleano--------")
	var boleano bool
	fmt.Println(boleano)

	fmt.Println("\n-------Erro--------")
	var erro error = errors.New("Erro interno")
	fmt.Println(erro, "\n ")
}