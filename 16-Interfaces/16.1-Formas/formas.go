package main

import (
	"fmt"
	"math"
)

type Forma interface { //Interfaces só possuem assinatura de métodos
	area() float64
}

func escreverArea(forma Forma) {
	//fmt.Printf("A área da forma é %0.2f u.a.\n", forma.area())
	forma.area()
}

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) area() (area float64){
	area = r.altura * r.largura
	fmt.Printf("A área do retangulo é %0.2f u.a.\n", area)
	return
}

func (c Circulo) area() (area float64) {
	area = math.Pi * math.Pow(c.raio, 2)
	fmt.Printf("A área do círculo é %0.2f u.a.\n", area)
	return
}

func main() {
	ret := Retangulo{10, 15}
	escreverArea(ret)

	circ := Circulo{10}
	escreverArea(circ)
}
