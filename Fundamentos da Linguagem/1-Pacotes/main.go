package main

import (
	"fmt"
	"modulo/Auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	// Função Escrever registra uma mensagem no CLI
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	// Pacote externo checkmail.ValidateFormat serve para validar um corpo de email e se o formato corresponde
	erro := checkmail.ValidateFormat("erro proposital")
	fmt.Println(erro)
}
