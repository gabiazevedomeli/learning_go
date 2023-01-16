package main

import "fmt"

/* Regras da linguagem:

- Não é permitido fazer operações matemáticas com variáveis de tipos diferentes. 
	Por exemplo: Não posso somar uma variável int16 + int32, dá erro porque elas são de tipos
	diferentes, se fossem duas variáveis int16 seria possível.

- Operadores relacionais retornam sempre um booleano, true / false.

- Operadores unários só funcionam para operações de incremento (+) ou decremento (-) de
	variáveis. O Go só aceita operadores unários pós fixados. 
	Exemplo: numero++ ou numero--

- Não existem operadores ternários no Go.
	Utiliza-se if/else para condicionais.

*/

func main() {
	// Aritméticos

	sum := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	mod := 10 % 2

	fmt.Println(sum, sub, div, mult, mod)

	var number1 int16 = 10
	var number2 int16 = 25

	som2 := number1 + number2
	fmt.Println(som2)

	// Operadores de Atribuição

	fmt.Println("----------------------------------------------------------------")
	var variable1 string = "String"
	variable2 := "String2"

	fmt.Println(variable1, variable2)

	// Operadores Relacionais

	fmt.Println("----------------------------------------------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos

	fmt.Println("----------------------------------------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores Unários - para incremento do valor de uma variável

	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 10 // numero = numero / 3
	numero %= 3 // numero = numero % 3
	fmt.Println(numero)

	// Operadores ternários (condicionais simples)

	fmt.Println("----------------------------------------------------------------")

	var text string

	if numero > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}

	fmt.Println(text)

}