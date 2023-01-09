package main

import "fmt"

/* Go é fortemente tipado
Existem duas maneiras de declarar variáveis em Go:
1 - Deixando o tipo dela implícito (através da inferência de tipos)
2 - Deixando o tipo dela explícito

Todas as variáveis serão tipadas, mas não necessariamente eu preciso fazer isso manualmente.
*/ 

/* Go não compila com delcarações de imports e variáveis que não foram usadas */

func main() {
	var variable1 string = "Gabriela" // declaração explícita
	fmt.Println(variable1)

	variable2 := "Mercado Pago" // declaração implícita
	fmt.Println(variable2)

	// declarando múltiplas variáveis ao mesmo tempo explicitamente

	var (
		variable3 string = "Mercado Libre"
		variable4 string = "Proximity Marcketplace"
	)

	fmt.Println(variable3, variable4)

	// declarando múltiplas variáveis ao mesmo tempo implicitamente

	variable5, variable6 := "Buyer", "Seller"
	fmt.Println(variable5, variable6)

	// declarando constantes

	const constant1 string = "Constantes em GO"
	fmt.Println(constant1)

	// caso eu queria substituir valores de duas variáveis 

	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
	
}