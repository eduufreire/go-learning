package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter you city: ")
	city, _ := reader.ReadString('\n')
	fmt.Print("You live in " + city)

	fmt.Print("Enter you name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Name: ", name)

	fmt.Print("Enter any number: ")
	number, _ := reader.ReadString('\r')
	convertedNumber, err := strconv.Atoi(number)
	fmt.Print(err)
	if convertedNumber >= 1 && convertedNumber <= 10 {
		fmt.Print("Seu numero está entre 1 e 10")
	} else {
		fmt.Print("Seu número é maior que 10 ou menor que 1")
	}
}
