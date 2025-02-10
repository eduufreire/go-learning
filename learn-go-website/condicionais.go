package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func pow(x, n, lim float64) float64 {
	// if pode ter uma declaracao antes da execucao da condicao
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func usandoSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func dataSwitch() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}

func switchSemCondicao() {
	// um switch sem condicao é o mesmo que true
	// maneira para fazer uma cadeira de if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default: 
		fmt.Println("Good evening")
	}
}


func switchExamples() {
	// podemos usar virgula para colocar mais de uma expressao no case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}


func usandoDefer() {
	// adia a execucao de uma funcao ate o final do retorno da funcao
	defer fmt.Println("world")
	fmt.Println("Hello")

	// chamadas de funcoes adiadas sao armazenadas em uma pilha
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}	


func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	usandoSwitch()
	dataSwitch()
	switchSemCondicao()
	usandoDefer()
}