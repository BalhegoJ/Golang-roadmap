package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

//Em Golang uma herança funciona passando o tipo de uma struct dentro da struct que queremos herdar os atributos como exemplificado abaixo, a struct pessoa sendo invocado na struct estudante
type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main (){
	
	//Na instanciação de variavel com tipo criado (struct) com herança, invocamos a herança na definição dos atribuotos
	pessoa1 := pessoa{"João Victor", "Balhego", 25, 180}
	estudante1 := estudante{pessoa1, "Ciencia da computação", "Univali"} 

	fmt.Println(pessoa1)
	fmt.Println(estudante1)

	//Para acessar os valores de cada instancia, funciona normalmente
	fmt.Println("nome: ", estudante1.nome, "\nSobrenome: ", estudante1.sobrenome, "\nidade: ", estudante1.idade, "\nAltura em cm: ", estudante1.altura, "\nFaculdade: ", estudante1.faculdade, "\nCurso: ", estudante1.curso)

}