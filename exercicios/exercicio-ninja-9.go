package main

import "fmt"

const (
	_ = 0
	a = iota + 2025
	b 
	c 
	d
)

func main() {

	// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	fmt.Print(a, b, c, d)


}
