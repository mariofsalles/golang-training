package main

import "fmt"

func main(){

	// Aritmeticos
	fmt.Println("***Aritmeticos")
	var soma int8 = 1 + 2
	fmt.Println(soma)
	subtracao := 5 - 3
	multiplicacao := 5 * 3
	divisao := 5 / 3
	//restoDivisao := 5 % 3
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(5 % 3)

	// Atribuição
	fmt.Println("\n***Atribuição")
	var var1 string = "texto 1"
	fmt.Println(var1)
	var2 := "texto 2"
	fmt.Println(var2)

	//Relacionais
	fmt.Println("\n***Relacionais")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)
	
	// Logicos
	fmt.Println("\n***Logicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	
	// Unários, não exitem ternários
	fmt.Println("\n***Unários")
	numeroMais := 10
	numeroMais ++
	
	numeroMenos := 14
	numeroMenos --
	
	fmt.Println(numeroMais)
	fmt.Println(numeroMenos)
	

}