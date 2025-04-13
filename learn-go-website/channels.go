package main

import "fmt"

func main() {
	// cria um canal e diz qual tipo sera transmitido nele
	ch := make(chan bool)
	go numbers1(ch)
	<-ch
	go letters(ch)
	<-ch
	fmt.Println("\nFim da execução")
}

// definindo uma funcao que escrevera valores em um canal do tipo inteiro
func numbers1(done chan<- bool) {
	for i := range 10 {
		fmt.Print(i, " ")
	}
	done <- true
}

// difinindo uma funcao que lerá valores de um canal do tipo inteior
func letters(done chan<- bool) {
	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("%c ", i)
	}
	done <- true
}