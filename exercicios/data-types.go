package main

import "fmt"

// https://github.com/bregman-arie/go-exercises/blob/main/exercises/data_types/exercise.md
func main() {
	var food = "Pizza"
    var slices = 4
    var pineappleOnPizza = true

	fmt.Printf("%T %T %T", food, slices, pineappleOnPizza)
}

// resposta: dara erro pois nao foi definido o tipo da variavel 'userName'
// func main() {
//     var userName
//     userName = "user"
//     fmt.Println(userName)
// }

// resposta: dará certo pois o Go irá inferir o tipo da variável
// func main() {
//     var userName = "user"
//     fmt.Println(userName)
// }

