package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 14

	fmt.Println("map: ", m)

	// se o valor n existir, é retornado um zero-value
	// de acordo com o tipo do map
	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	delete(m, "k2")
	fmt.Println("map: ", m)

	// remove tudo
	clear(m)

	// retorna um segundo valor, dizendo se existe a chave ou nao
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n === n2")
	}
	
}
