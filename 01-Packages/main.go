package main

/* funções escritas com letra maiúscula são visíveis em outros pacotes (públicas), 
já funções que começam com letra minúscula são privadas, 
só podem ser chamadas dentro do mesmo pacote em que foram criadas. */

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("gabmeli@gmail.com")
	fmt.Println(erro)
}