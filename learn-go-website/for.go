// go só tem o for
package main

import "fmt"

func main() {
	sum := 0
	// a instrucao inial e a pos-instrucao sao opcionais
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for ; sum < 1000 ; {
		sum += 2
	}
	fmt.Println(sum)

	//for é o "while" de go
	teste := 1
	for teste < 1000 {
		teste += teste
	}
	fmt.Println(teste)

	nums := []int{1, 2, 3, 4, 5}
	for key, value := range nums {
		fmt.Println(key, value)
	}

}