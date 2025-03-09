package main

import "fmt"

// o * na frente do tipo indica que receberemos o endereço
// de memória daquele tipo
func dobrar(value *int) {
	*value *= 2
}

func trocar(a, b *int) {
	*b, *a = *a, *b
}

func novoValor() *int {
	valor := 10
	fmt.Println(&valor)
	return &valor
}

type Pessoa struct {
	Nome string
	Idade int8
}

func envelhecer(p *Pessoa) {
	p.Idade++
}

func incrementarArray(arr *[5]int) {
	for i := range arr {
		arr[i]++
	}
}

func main() {
	
	x := 10
	fmt.Println("Antes ", x)
	// passando o endereço
	dobrar(&x)
	fmt.Println("Depois ", x)


	z, a := 10, 20
    fmt.Println("Antes:", z, a)
    trocar(&z, &a)
    fmt.Println("Depois:", z, a)

	p := novoValor()
	fmt.Println("Valor: ", *p)

	pessoa := Pessoa{"Joao", 25}
	fmt.Println("Antes:", &pessoa)
	envelhecer(&pessoa)
	fmt.Println("Depois:", pessoa)

	fmt.Print("\n\n")
	nums := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Antes:", nums)
    incrementarArray(&nums)
    fmt.Println("Depois:", nums)
}


