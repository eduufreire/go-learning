package main

import (
	"fmt"
)

func main() {

	customers := 4
	packages := 100
	delivered := 20

	fmt.Printf("I have %v packages to deliver \n", packages)

	packagesRemaining := packages - delivered
	fmt.Printf("I'm delivered %v. Now, remaining %v packages \n", delivered, packagesRemaining)

	fmt.Printf("I'm distributed equally %v packages between %v customers \n", packagesRemaining / customers, customers)



}