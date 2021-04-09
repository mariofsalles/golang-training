package main

import "fmt"

var f = func() {
	fmt.Println("Funcao")
}

var soma = func(a int8, b int8) {
	fmt.Println(a + b)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := (n1 + n2)
	subtracao := n1 - n2
	return soma, subtracao
}


func main() {
	soma(2, 5)
	f()

	_ , resultadoSubtracao := calculosMatematicos(5, 3)
	fmt.Println(resultadoSubtracao)
}
