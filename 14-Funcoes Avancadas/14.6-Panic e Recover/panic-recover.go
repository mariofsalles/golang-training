package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar execução")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6!!!")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}
