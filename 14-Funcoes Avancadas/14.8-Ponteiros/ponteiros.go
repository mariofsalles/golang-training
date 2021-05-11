package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * (-1)
}

func inverteSinalPonteiro(numero *int) {
	*numero = (*numero) * (*numero) // elevei ao quadrado
}

func main() {
	numero1 := 20
	invertido1 := inverteSinal(numero1) // houve copia do valor da variavel e posto nesta
	fmt.Println(invertido1)
	fmt.Println(numero1) // observa-se que não ocorreu alteração no valor

	numero2 := 40
	fmt.Println(numero2)  // a variável numero 2 aqui está intacta
	fmt.Println(&numero2) // este é o endereço de memoria dele

	fmt.Println("---------------")
	inverteSinalPonteiro(&numero2) // agora ela foi alterada, pois passou o endereço de memória &numero2 para efetuar a operação
	fmt.Println(numero2, &numero2) // ocorreu alteração no valor, mas o endereço de memoria não se altera como se espera

}
