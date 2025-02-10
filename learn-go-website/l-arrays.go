package main

import (
	"fmt"
)

func main() {

	arr1 := [5]int64{1, 2, 3, 4, 5}
	for i := range len(arr1) {
		fmt.Println(arr1[i])
	}

	var arr2 [5]int
	fmt.Println(arr2[0])

	arr3 := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println("len ", len(arr3))

	// podemos dizer quais/ate qual indice sera inicializado com valor 0
	// podemos dizer que o indice 4 ja começara com 400
	arr4 := [...]int{100, 4: 400, 500}
	fmt.Println(arr4)

	var matriz1 [2][3]int
	fmt.Println(matriz1)

	matriz2 := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(matriz2)


	keyValue := [5]string {3: "Chris", 4: "Ron"}
	fmt.Println(keyValue)

	tenNumbers := [...]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("len ", len(tenNumbers), " values: ", tenNumbers)

	arrayNames := [5]string {"edu", "jose", "maria", "joao", "lucia"}
	fmt.Println("len ", len(arrayNames), " values: ", arrayNames)


}
