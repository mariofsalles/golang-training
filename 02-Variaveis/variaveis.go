package main

import "fmt"

func main(){
	var variavel1 string = "pato" //explicita
	variavel2 := "galinha" //implicita (inferencia de tipos)
	fmt.Println("variavel1:",variavel1,"\nvariavel2:",variavel2)
	
	var (
		variavel3 string = "casa"
		variavel4 string = "cachorro"
	)

	variavel5, variavel6 := "macaco", "leao"
	
	fmt.Println("variavel3:",variavel3,"\tvariavel4:",variavel4)
	fmt.Println("variavel5:",variavel5)
	fmt.Println("variavel6:",variavel6)
	
	const contante1 string = "valor constante"
	fmt.Println("contante1:",contante1)
	
}