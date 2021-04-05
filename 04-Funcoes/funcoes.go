package main

import "fmt"

var f = func() {
	fmt.Println("Funcao")
}

var soma = func(a int8, b int8) {
	fmt.Println(a + b) 
}


func main(){
	soma(2,5)
	f()
}