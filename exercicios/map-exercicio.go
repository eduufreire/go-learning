package main

import (
	"fmt"
	"strings"
)

func main() {

	texto := "go é rápido go é simples"
	palavras := strings.Split(texto, " ")
	resultado := make(map[string]int) 
	for i := range len(palavras) {

		_, check := resultado[palavras[i]]

		if !check {
			resultado[palavras[i]] = 1
		} else {
			resultado[palavras[i]] += 1
		}

		
	}
	fmt.Println(resultado)

}
