package main

import (
	"fmt"
)

func main() {

	var a, b, c float64 = 6.0, 4.0, 2.1

	teste1 := a+b > c
	teste2 := a+c > b
	teste3 := b+c > a

	perimetro := 0.0
	area := 0.0
	if teste1 && teste2 && teste3 {
		perimetro = a + b + c
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else {
		area = (a + b) * c / 2
		fmt.Printf("Area = %.1f\n", area)
	}

}
