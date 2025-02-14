package main

import (
	"fmt"
)

func main() {

	// Crie constantes tipadas e não-tipadas.
	// Demonstre seus valores.

	const NOME string = "Eduardo"
	const IDADE = 20
	const MAIOR_IDADE = IDADE >= 18
	const CIDADE string = "São Paulo"

	fmt.Printf("%v %T\n", NOME, NOME)
	fmt.Printf("%v %T\n", IDADE, IDADE)
	fmt.Printf("%v %T\n", MAIOR_IDADE, MAIOR_IDADE)
	fmt.Printf("%v %T\n", CIDADE, CIDADE)

}
