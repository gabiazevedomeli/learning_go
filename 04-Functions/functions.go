package main

import "fmt"
/* Funções:

- Declara-se utilizando a palavra reservada func;
- Pode receber ou não parâmetros;
- É necessário tipar os parâmetros recebidos;
- Caso a função tenha cláusula de retorno o mesmo é especificado (somente o tipo de dado do retorno) 
	após delcaração dos parâmetros da função, ou seja, após o parêntese de fechamento dos parâmetros.

- Funções em GO podem ter mais de um retorno (???). Sim, é possível, delcarando quais e quantos retornos
	serão programados para cada função, basta adicionar mais de um retorno ou vários, entre parênteses, após
	delcaração dos parâmetros.
	- Ao chamar funções com mais de um retorno é necessário passar na chamada a quantidade de variáveis
		de acordo com o número de retornos, cara variável abrigará o resultado de cada retorno.

	- Quando eu quero chamar funções com mais de um retorno e na chamada eu não preciso de todos os 
		retornos basta suprimir o(os) retornos indesejados com um underline (_). É necessário atentar-se com
		a ordem de parâmetros e retornos.
	Exemplo abaixo na função mathCalculations()

- Se eu quiser declarar mais de um parâmetro do mesmo tipo em uma função posso tipar apenas o último 
	parâmetro e o primeiro terá seu tipo inferido de acordo com a delcaração explícita de tipo do último parâmetro.
	exemplo abaixo na função mathCalculations()

*/
func sum(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func mathCalculations(num1, num2 int8) (int8, int8) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main() {
	sum := sum(10, 20)
	fmt.Println(sum)

	var function = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	functionResult := function("Function text")
	fmt.Println(functionResult)

	/* Chamando funções com mais de um retorno */ 
	resultsSum, resultsSub := mathCalculations(10, 15)
	fmt.Println(resultsSum, resultsSub)

	/* Chamando funções com mais de um retorno, suprimindo retornos indesejados */
	resultsSum1, _ := mathCalculations(10, 15)
	fmt.Println(resultsSum1)
}
