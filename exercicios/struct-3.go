package main

import "fmt"

type Produto struct {
	Nome   string
	Codigo int
	Preco  float64
}

func (p *Produto) Apresentar() {
	fmt.Printf("Código: %d | Produto: %s | Preço: %.2f\n", p.Codigo, p.Nome, p.Preco)
}

func main() {

	produtos := []Produto{
		Produto{
			Nome:   "Sabontete",
			Codigo: 1,
			Preco:  3.99,
		},
		Produto{
			Nome:   "Macarrão",
			Codigo: 2,
			Preco:  4.59,
		},
		Produto{
			Nome:   "Refrigerante",
			Codigo: 3,
			Preco:  11.67,
		},
	}

	indiceMaiorPreco := 0
	maiorPreco := 0.0
	fmt.Println("Produtos: ")
	for key, produto := range produtos {
		if produto.Preco > maiorPreco {
			indiceMaiorPreco = key
		}
		produto.Apresentar();
	}

	fmt.Println("\nProduto com maior preço: ")
	produtos[indiceMaiorPreco].Apresentar()

}