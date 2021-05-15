package main

import (
	"bytes"
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
	cachorro1 := Cachorro{"Rex","Dalmata",3}
	fmt.Println(cachorro1)

	cachorro1EmBytes, erro := json.Marshal(cachorro1)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro1EmBytes)

	cachorro1EmJSON := bytes.NewBuffer(cachorro1EmBytes)
	fmt.Println(cachorro1EmJSON)

	fmt.Println("*************************")

	cachorro2 := map[string]string{
		"nome":"Totó",
		"raca":"Pastor Alemão",
		"idade_cachorro": "4",
	}
	cachorro2EmBytes, erro := json.Marshal(cachorro2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmBytes)

	cachorro2EmJSON := bytes.NewBuffer(cachorro2EmBytes)
	fmt.Println(cachorro2EmJSON)
}
