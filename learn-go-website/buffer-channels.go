package main

import "fmt"

func main() {

	cn := make(chan int, 3);
	go numbers2(cn)

	for v := range cn {
		fmt.Println("Leitura: ", v)
	}
}

func numbers2(ch chan<- int) {
	for i := range 10 {
		fmt.Println("Escrita: ", i)
		ch <- i
	}	
	fmt.Println("Fim da escrita")
	close(ch)
}