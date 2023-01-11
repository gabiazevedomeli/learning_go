package main

import (
	"errors"
	"fmt"
)

/* Há dois tipos de números no GO:

1 - Inteiros:

Tipos existentes:

int, int8, int16, int32, int64
Cada número nesses inteiros, 8, 16, 32, 64
corresponde ao número de bits que o número inteiro ocupa na memória

int - quando colocamos int, sem especificar os bits, o go utiliza a arquitetura do computador para
configurar a variável. Se o computador é 64bits ele vai criar um int64 se for 32 cria um int32.

se declararmos a variável através de inferência de tipo o go seta o int de acordo com a
arquitetura do computador.

Suporta números negativos.

uint, uint8, uint16, uint32, uint64 - unsygned int
São inteiros somente positivos, sem sinal.

2 - Reais (float):

Números com ponto flutuante. Existem apenas dois tipos.

Tipos existentes:

float32, float64

Para declarar com inferência de tipos basta declarar sem colocar explicitamente o tipo
e a linguagem infere o tipo de acordo com a arquitetura do computador (32 ou 64)

3 - Strings:

Sempre declarar usando aspas duplas ("")

Quando declaramos um caracter usando aspas simples ('')
teremos um valor numérico, correspondente ao número do caracter na
tabela ASC.

4 - Booleans:

Tipos existentes:

true, false

5 - Erros:

Tipos existentes:

error

Valor default - <nil> (entende-se null)
Valor atribuído - usa-se o pacote errors do go para tratamento de erros.

*/

/* ---------------------------------------------------------------- */

/* Conceitos da linguagem:

- Valor 0 ou valores default:

É o valor que será atribuído a uma variável quando esta não é inicializada em sua declaração.

Para strings o valor 0 é uma string vazia ("")
Para números inteiros ou reais é 0
Para booleans é false
Para erro é nil (valor null)

- Ponto e vírgula:

Não é necessário colocar explicitamente o ponto e vírgula
para delimitar uma ação para o compilador da linguagem.
O próprio compilador no momento da compilação adiciona o ;

- Declarações implícitas não aceitam variáveis não inicializadas, portanto,
é necessário inicializá-la para os casos de delcaração implícita de variável em go.
*/

func main() {

	/* Números inteiros */

	number := 10000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	// alias
	// int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	// byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	/* Números reais (de ponto flutuante) */

	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 1230000000000.45
	fmt.Println(realNumber2)

	realNumber3 := 12345.67
	fmt.Println(realNumber3)

	/* Strings e char */

	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	/* Booleanos */

	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolena2 bool
	fmt.Println(boolena2)

	/* Erro */

	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Internal error")
	fmt.Println(erro2)
}