package main

import (
	"fmt"
)

func main() {
	// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
	number := 10

	fmt.Printf("%b\n", number)
	fmt.Printf("%d\n", number)
	fmt.Printf("%#x\n", number)

}