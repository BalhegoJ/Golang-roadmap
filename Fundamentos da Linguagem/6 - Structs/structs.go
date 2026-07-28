package main

import "fmt"

//Tipo de dado que agrupa varias variaveis em uma estrutura, parecido com um exemplo com um objeto em POO
type usuario struct {
type usuario struct {
	nome string
	idade int8
	endereço endereco
}

//Podemos usar uma struct dentro de outra struct, como por exemplo a struct endereço abaixo sendo usada como um campo na struct usuário acima
type endereco struct{
	rua string
	numero int16
	complemento string
}

func main(){

	var usuario1 usuario

	//Podemos atrubuir um campo e valor por vez
	usuario1.idade = 18
	usuario1.nome = "João Victor Balhego"
	fmt.Println(usuario1)

	//Podemos atrubuir todos os valores da struct de uma vez só, seguindo a ordem de declaração das varaiveis
	usuario2 := usuario{"Fulano", 21, endereco{"Ruinha da silva", 200, "Centro Historico"}}
	fmt.Println(usuario2)

	//Podemos atrubuir somente um valor na instancia de uma variavel da struct
	usuario3 := usuario{nome: "Ciclano"}
	fmt.Println(usuario3)


}
