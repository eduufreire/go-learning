package main

import (
	"fmt"
)

func main() {

	sliceOne := []int{}
	fmt.Println(sliceOne)

	// cap() capacidade do slice
	// len() quantidade de elementos do slice
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	// criando a partir de um array
	// slice := array[start:end] end == index - 1
	myArray := [5]int{1, 2, 3, 4, 5}
	sliceTwo := myArray[1:3]
	fmt.Println(sliceTwo)

	// make()
	sliceThree := make([]int, 3)
	fmt.Println(len(sliceThree))
	fmt.Println(cap(sliceThree))
	fmt.Println(sliceThree)

	// append
	sliceAppended := append(sliceTwo, 25, 26, 27)
	fmt.Println(sliceAppended)

	slicesAppend := append(sliceOne, sliceThree...)
	fmt.Println(slicesAppend)

	// copy
	// cria um array com somente os elementos requeridos no slice, economizando memoria
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	neededNumbers := numbers[:len(numbers)-5]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Println(numbersCopy)
	fmt.Println(cap(numbersCopy))

}
