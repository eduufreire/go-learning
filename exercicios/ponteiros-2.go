package main

import (
	"fmt"
)

func criarNumero() *int {
	valor := 100
	return &valor
}

type ContaBancaria struct {
	Titular string
	Saldo float64
}

func depositar(cb *ContaBancaria, valor float64) {
	cb.Saldo += valor
}

func modificarElemento(sl *[]int) {
	(*sl)[0] = 100
}

func soma(a, b int) int {
	return a + b
}

func main() {

	ponteiroValor := criarNumero()
	fmt.Println(*ponteiroValor)

	conta := ContaBancaria{
		Titular: "Edu",
		Saldo: 10.0,
	}
	depositar(&conta, 10)
	fmt.Println(conta.Saldo)

	mySlice1 := []int{1, 2, 3, 4}
	fmt.Println(mySlice1)
	modificarElemento(&mySlice1)
	fmt.Println(mySlice1)
	
	myMap1 := make(map[string]int)
	pointerMap := &myMap1

	(*pointerMap)["teste"] = 1
	(*pointerMap)["teste1"] = 2
	(*pointerMap)["teste2"] = 3

	fmt.Println(*pointerMap)

}