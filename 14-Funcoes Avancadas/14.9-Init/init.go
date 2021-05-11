package main

import "fmt"

var n int

func init()  {
	fmt.Println("Executa antes da main: init")
	n = 10
}

func main()  {
	fmt.Println("Função main está sendo executada")
	fmt.Println(n)
}