package main

import "fmt"

func main() {
	// funcao anonimas
	multiplica := func(x int) int {
		return x * 2
	}

	resultado := multiplica(2)
	fmt.Println(resultado)

	fmt.Print(plus(1, 2))

	a, b := multiplyValues()
	fmt.Print(a, b)
	_, c := multiplyValues()
	fmt.Print(c)

	sum(1, 2)
	sum(1, 3, 5, 7)
	sum(2, 4, 6, 8, 10)
}

// é preciso ter o retorno explicito
func plus(a, b int) int {
	return a + b
}

// multiplo valores de retorno
func multiplyValues() (int, int) {
	return 3, 7
}

// funcoes variadicas
// recebe n argumentos
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Print(total)
}
