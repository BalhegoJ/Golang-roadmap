package main

import "fmt"

func main() {
	//As variaveis em Go lang podem ser apresentadas de duas maneiras, explicitas e implicitas
	var variavel1 string = "Variavel explicita"
	fmt.Println(variavel1)

	variavel2 := "variavel implicita"
	fmt.Println(variavel2)

	//É possivel declara uma cadeia de variaveis
	var (
		variavel3 uint8 = 10
		variavel4 int8 = -10
	)
	fmt.Println(variavel3, variavel4)

	//É possivel declarar mais de uma variavel por ver
	variavel5, variavel6 := "Primeira", "segunda"
	fmt.Println(variavel5, variavel6)

}
