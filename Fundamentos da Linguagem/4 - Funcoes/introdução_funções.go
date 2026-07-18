package main

import (
	"fmt"
)

//O grande diferencial em Golang é que podemos ter mais de um retorno em uma função
//Apos a decalaração de parametros, também declaramos o tipo e quantidade de retornos (No exemplo são 2)

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func calculosMat(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {

	//Podemos atribuir a uma variavel a função e os parametros, o retorno da função será o valor que a variavel receberá
	soma := somar(5, 15)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Texto da função F")

	//Nessa função teremos 2 retornos que são a soma e subtração dos parametros passados, já que ela retorna mais de um valor
	fmt.Println(calculosMat(10, 5))

	//Podemos atribuir um retorno apenas de uma função com mais de um retorno com o caractere "_"
	//Ao atribuir os retornos da função em variaveis, o caractere "_" que seria uma das variaveis que receberia o valor do retorno, indica que o Go pode ignorar aquele retorno, importante saber a ordem dos fatores na função
	resultadoSoma, _ := calculosMat(30, 12)
	println(resultadoSoma)
}
