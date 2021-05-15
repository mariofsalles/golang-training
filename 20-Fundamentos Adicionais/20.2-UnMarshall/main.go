package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade_cachorro"`
}

func main() {
	
	cachorro3JSON := `{"nome":"Rex","raca":"Dalmata","idade_cachorro":3}`
	//cachorro3 := Cachorro{}, ou o valor abaixo
	var cachorro3 Cachorro
	if erro := json.Unmarshal([]byte(cachorro3JSON), &cachorro3); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro3)
		
	cachorro4JSON := `{"nome":"Toto","raca":"Pastor Alemão"}`
	cachorro4 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro4JSON), &cachorro4); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro4)

}
