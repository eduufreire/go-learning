package main

import (
	"fmt"
)

// https://github.com/bregman-arie/go-exercises/blob/main/exercises/variables/exercise.md
func main() {

	name, age := "Eduardo", 20
	fmt.Printf("My name is %v and I'm %v years old", name, age)

}


// func questionOne() {
// 	// nao sera compilado pois é declarada e inicializada uma variavel objective
// 	// porem ela nao é utilizado. 
// 	// Em Go, tudo deve que é declarado, deve ser utilizado
// 	fmt.Println("Skynet Beta Testing")

//     var objective="Defend Humanity"
// }

// func questionTwo() {
// 	// o código nao ira compilar pois estamos tentando atribuir um novo valor para squares
// 	// isso nao é possivel pois ela é uma constante
// 	const squares=2
//     var circles=0
    
//     fmt.Println("Squares:", squares)
//     fmt.Println("Circles:", circles)

//     circles=7

//     fmt.Println("Squares:", squares)
//     fmt.Println("Circles:", circles)
// }