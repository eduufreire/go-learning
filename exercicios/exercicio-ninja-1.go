package main

import (
	"fmt"
)

func main() {
	x, y, z := 42, "James Bond", true
	fmt.Println(x, y, z)

	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
	fmt.Print(z, "\n")
}