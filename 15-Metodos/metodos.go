package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
}

func (user Usuario) salvar() { // o metodo se chama salvar e está dentro da estrutura de Usuário, dessa forma podemos colocar os parametros do Usuario
	fmt.Printf("Salvando os dados do usuario '%s' no banco de dados, ele possui '%d' anos\n", user.nome, user.idade)
}

func (user Usuario) maiorDeIdade() (testeIdade bool) {
	// return user.idade >= 18 --> neste caso o retorno não precisa ser nomeado
	fmt.Println(user.idade >= 18)
	return 
}

func (user *Usuario) fazerAniversario(){
	user.idade++
}

func main() {

	usuario1 := Usuario{"Fulaninho", 20}
	usuario1.salvar()

	usuario2 := Usuario{"Beltrano", 17}
	usuario2.salvar()
	// maiorDeIdade := 
	usuario2.maiorDeIdade()
	//fmt.Println(maiorDeIdade) 
	
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

	//maiorDeIdade2 := 
	usuario2.maiorDeIdade()
	//fmt.Println(maiorDeIdade2)
}
