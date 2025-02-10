package main

import (
	"fmt"
)

type meuTipo int
var x meuTipo
var y int

func main() {

	fmt.Printf("%v type: %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v type: %T\n", y, y)


}