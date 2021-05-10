package main

import (
	"fmt"
	"time"
)

func main()  {
	// CONCORRENCIA (PODE EXECUTAR AO MESMO TEMPO, MAS O PRINCIPAL É QUE HÁ REVESAMENTO DE TAREFAS) != PARALELISMO (TAREFAS EXECUTADA AO MESMO TEMPO)
	go escrever("Olá Mundo!")
	escrever("Programando em Go!")

}

func escrever(texto string){
	for{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}