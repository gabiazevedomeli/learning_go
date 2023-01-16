package main

import "fmt"

/* Structs

- São coleções de campos, cada campo tem um nome e um tipo. Quando temos vários campos,
	várias informações relativas a uma mesma coisa, usamos struct.
	O struct é uma estrutura que lembra a estrutura de classes no Java. É criado um tipo específico e dentro dele
	atribuimos as variáveis e seus respectivos tipos.

- Quando instanciamos um struct sem passar nenhum valor, ele infere os valores default,
	de acordo com os tipos das variáveis declaradas dentro do struct.

*/

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u user
	u.name = "Davi"
	u.age = 21
	fmt.Println(u)

	address1 := address{
		"Rua dos bobos", 0,
	}

	user2 := user{
		"Gabriela", 22,
		address1,
	}

	fmt.Println(user2)

	user3 := user{name: "Heitor"}
	fmt.Println(user3)

}