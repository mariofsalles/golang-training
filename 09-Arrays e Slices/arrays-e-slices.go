package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Array e Slices")

	var array1[5]int
	array1[0]=1
	fmt.Println(array1)
	
	array2 := [5]string{"pos1","pos2","pos3","pos4","pos5"}
	fmt.Println(array2)
	
	array3 := [...]int{1, 2, 3, 4, 5} //não deixa o array com tamanho dinamico
	fmt.Println(array3)
	fmt.Println(len(array3))

	slice1 := []int{1,2,3,4,5,6}
	fmt.Println(slice1)
	
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))
	
	slice1 = append(slice1, slice1[5])
	fmt.Println(slice1)

	slice2 := array2[1:3] //inclui o elemento do 1º indice(inclusivo) e não inclui o elemento do último indice(exclusivo)
	fmt.Println(slice2)
	
	array2[1] = "posição alterada" 
	fmt.Println(slice2) // funciona da mesma forma que um ponteiro, ou seja essa alteração mesmo que neste ponto reflete na estrutura inicial
	
	fmt.Println("-------------")
	// Arrays internos
	
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println("tamanho: ", len(slice3))
	fmt.Println("capacidade: ", cap(slice3))
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	fmt.Println("tamanho: ", len(slice3))
	fmt.Println("capacidade: ", cap(slice3))
	fmt.Println(slice3)


}
