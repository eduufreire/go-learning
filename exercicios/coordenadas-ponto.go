package main

import (
	"fmt"
)

func main() {
	x, y := -1.0, -0.1

	origem := (x == y && y == 0)
	eixoY := (x == 0 && !origem)
	eixoX := (y == 0 && !origem)

	q1 := (x > 0 && y > 0)
	q2 := (x < 0 && y > 0)
	q3 := (x < 0 && y < 0)
	q4 := (x > 0 && y < 0)

	switch {
	case origem:
		fmt.Println("Origem")
	case eixoY:
		fmt.Println("Eixo Y")
	case eixoX:
		fmt.Println("Eixo X")
	case q1:
		fmt.Println("Q1")
	case q2:
		fmt.Println("Q2")
	case q3:
		fmt.Println("Q3")
	case q4:
		fmt.Println("Q4")
	default:
		fmt.Println("Not found")
	}
}
