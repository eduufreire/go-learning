package main

import (
	"fmt"
)

func main() {

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for range 2 {
		fmt.Println("oi")
	}


}