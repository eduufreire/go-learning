package main

import "fmt"

// composicao de tipos
// usados para agrupar dados
// ""parecido com uma class"""
// podemos usar struct dentro de struct

// struct anonimo
// struct usado para apenas um valor
type Cliente struct {
	Nome     string
	Idade    int
	Endereco string
	Email    string
}

// métodos
func (c Cliente) Apresentar() {
	fmt.Println("ola cliente tal tal tal bla bla bla")
}

func main() {
	cliente1 := Cliente{
		Nome:     "edu",
		Idade:    20,
		Endereco: "Rua",
		Email:    "edu@gmail.com",
	}

	fmt.Print(cliente1.Email)
	cliente1.Apresentar()

	// struct anonimo
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Print(dog)
}
