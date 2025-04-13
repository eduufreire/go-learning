package main

import (
	"fmt"
	"time"
)

func firstGoroutine() {
	fmt.Println("My goroutine")
}


func words() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 200)
	}
}

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go firstGoroutine()
	go numbers()
	go words()
	time.Sleep(time.Second * 1)
	fmt.Println("Minha funcao principal")
}