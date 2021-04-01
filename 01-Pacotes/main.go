package main

import (
	"fmt"
	"module01/auxiliar"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrecendo num arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("mariosalles@gmail.com")
	fmt.Println(erro)
}