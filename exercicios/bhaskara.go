package main

import (
	"fmt"
	"math"
)

func main() {

	/**
	 * Escreva a sua solução aqui
	 * Code your solution here
	 * Escriba su solución aquí
	 */
	var a, b, c = 0.0, 20.0, 5.0

	delta := (math.Pow(b, 2)) - (4 * a * c)

	fmt.Println(delta)

	if a == 0.0 || delta < 0 {
		fmt.Println("Impossivel calcular")
	} else {
		sqrtDelta := math.Sqrt(delta)
		r1 := (-b + sqrtDelta) / (2 * a)
		r2 := (-b - sqrtDelta) / (2 * a)
		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f", r2)
	}

}
