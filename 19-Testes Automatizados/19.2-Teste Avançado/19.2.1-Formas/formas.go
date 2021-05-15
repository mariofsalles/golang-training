package formas

import (
	"math"
)

//Interfaces só possuem assinatura de métodos
type Forma interface { 
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() (area float64){
	area = r.Altura * r.Largura
	//fmt.Printf("A área do retangulo é %0.2f u.a.\n", area)
	return
}

func (c Circulo) Area() (area float64) {
	area = math.Pi * math.Pow(c.Raio, 2)
	//fmt.Printf("A área do círculo é %0.2f u.a.\n", area)
	return
}