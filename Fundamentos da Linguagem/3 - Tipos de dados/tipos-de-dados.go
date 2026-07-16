package main

import (
	"errors"
	"fmt"
)

func main (){

	//Em Golang temos a variavel int (Assume a arquitetura da CPU (64 ou 32 bits))
	//Os demais Ints são (int8, int16, int32, int64)
	var num int = 100000000
	fmt.Println(num)

	//rune = int32
	var num2 rune = 1234
	fmt.Println(num2)

	//byte = int8
	var num3 byte = 55
	fmt.Println(num3)

	//Numeros reais são representados pelo tipo primitivo float 32 e float 64
	var num4 float32 = 10.0000
	fmt.Println(num4)

	var num5 float64 = 10009999999999.5
	fmt.Println(num5)

	//Em Golang temos a variavel float (Assume a arquitetura da CPU (64 ou 32 bits))
	num6 := 543.21
	fmt.Println(num6)

	//Em Go não tem o tipo char (Caractere unico)
	//Se declararmos uma variavel com um valor de aspas simples, quando o programa for compilado será mostrado na tela do CLi o valorna tabela ASCII daquele caractere  
	//O valor assume o numero da tabela ASCII daquele caractere, logo a variavel se transforma automatico em INT
	char := 'J'
	fmt.Println(char)

	//Variabel logica, comportamento padrão
	var booleano bool = true
	fmt.Println(booleano)

	//Variavel error é o grande diferencial de Go
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	//O valor padrão para o tipo de dado string é uma string vazia
	//O valor padrão para o tipo de dado int é 0
	//O valor padrão para o tipo de dado erro é <nil> (nulo)
	//O valor padrão para o tipo de dado boolean false
}
