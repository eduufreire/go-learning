package main

import "fmt"

type Pessoa1 struct {
	Nome   string
	Idade  int
	Altura float64
}

func (p *Pessoa1) Apresentar() {
	fmt.Printf("Nome: %s | Idade: %d | Altura: %.2f", p.Nome, p.Idade, p.Altura)
}


func main() {

	var nome string
	var idade int
	var altura float64

	fmt.Println("Insira o nome: ")
	fmt.Scanln(&nome)

	fmt.Println("Insira a idade: ")
	fmt.Scanln(&idade)

	fmt.Println("Insira a altura: ")
	fmt.Scanln(&altura)

	pessoa1 := Pessoa1{
		Nome: nome,
		Idade: idade,
		Altura: altura,
	}

	pessoa1.Apresentar()
}