package main

import "fmt"

func main()  {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	
	// Nessa impressão o valor será passado por cópia
	fmt.Println(var1, var2)
	
	
	var2++	
	// Nessa impressão o valor será passado por cópia
	fmt.Println(var1, var2)
	//**********************************************//

	// Pensar em ponteiro como se fosse uma variável que salva o endereço de memória de alguma coisa
	// Ponteiro é uma referência de memória

	var var3 int //guarda um valor <do tipo> inteiro
	var pointer *int //guarda o endereço de memoria de um valor <do tipo> inteiro

	var3 = 100
	pointer = &var3

	fmt.Println(var3, pointer)
	var3 = 150
	fmt.Println(var3, *pointer)
	fmt.Println(pointer, *pointer)



}