package main

import "fmt"

func main() {
	
	x := 10

	// pegando o endereço da variavel x e guardando em y
	// agora y é um pointer, pois guarda o endereço de memoria
	y := &x

	// o * na frente do ponteiro permite acessar o valor dentro do endereço
	fmt.Println(y, *y)

}