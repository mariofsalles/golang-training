package main

import "fmt"

func main() {
	fmt.Println("Mpas")

	usuario1 := map[string]string{
		"nome":      "Mario",
		"sobrenome": "Salles",
	}
	
	fmt.Println(usuario1)
	fmt.Println("Nome: ", usuario1["nome"])
	fmt.Println("Sobrenome: ", usuario1["sobrenome"])
	
	usuario2 := map[string]map[string]string{
		"nome" : {
			"primeiro":"Mario",
			"segundo":"Cesar",
		},
		"sobrenome": {
			"primeiro":"Fernandez",
			"segundo":"Salles",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
	
	usuario2["signo"] = map[string]string{
		"nome" :"Virgem",
	}
	fmt.Println(usuario2)
	
	delete(usuario2["nome"], "primeiro")
	fmt.Println(usuario2)
	
}
