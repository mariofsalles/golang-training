package main

import "fmt"

func main() {
	func() {
		fmt.Println("sem parâmetro")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Com parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Com parâmetro")

	fmt.Println(retorno)
}
