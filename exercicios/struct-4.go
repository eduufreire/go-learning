package main

import "fmt"

type Livro struct {
	Título        string
	Autor         string
	AnoPublicacao int
}

func (l *Livro) Apresentar() {
	fmt.Printf("Titulo: %s | Autor: %s | Ano Publicacao: %d\n", l.Título, l.Autor, l.AnoPublicacao)
}

func main() {

	livros := make([]Livro, 2)
	maioresDoisMil := make([]int, 2)

	quantidadeLivros := 0
	fmt.Println("Quantos livros deseja cadastar: ")
	fmt.Scan(&quantidadeLivros)

	var titulo, autor string
	var anoPublicacao int
	for quantidadeLivros > 0 {

		fmt.Println("Título do livro: ")
		fmt.Scan(&titulo)

		fmt.Println("Autor do livro: ")
		fmt.Scan(&autor)

		fmt.Println("Ano de publicação do livro: ")
		fmt.Scan(&anoPublicacao)

		livros = append(livros, Livro{
			Título:        titulo,
			Autor:         autor,
			AnoPublicacao: anoPublicacao,
		})

		quantidadeLivros--
	}

	fmt.Println("Todos os livros: ")
	for key, value := range livros {
		if value.AnoPublicacao > 2000 {
			maioresDoisMil = append(maioresDoisMil, key)
		}
		if value.Título != "" {
			value.Apresentar()
		}
	}

	fmt.Println("\nLivros após os anos 2000: ")
	for _, value := range maioresDoisMil {
		if livros[value].Título != "" {
			livros[value].Apresentar()
		}
	}

}
