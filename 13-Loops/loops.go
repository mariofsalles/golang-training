package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	fmt.Println("************")

	i := 0
	for i < 10 {
		i++
		fmt.Printf("Incrementando %v\n", i)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("************")

	for j := 0; j < 10; j++ {
		fmt.Printf("Incrementando %v\n", j)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("************")

	nomes := [3]string{"mario", "clara", "luiz"}
	for _, nome := range nomes {
		fmt.Printf("nome: %s\n", nome)
		time.Sleep(200 * time.Millisecond)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, "-", string(letra))
		time.Sleep(200 * time.Millisecond)
	}

	usuario1 := map[string]string{
		"nome":      "Mario",
		"sobrenome": "Salles",
	}

	fmt.Println(usuario1)

	for chave, valor := range usuario1 {
		fmt.Println(chave, valor)
	}

	count := 1
	// igual a for true
	for {
		fmt.Println("Lopp infinito: ", count)
		time.Sleep(200 * time.Millisecond)
		count++
	}

}
