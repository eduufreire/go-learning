package main

import (
	"fmt"
)

// https://judge.beecrowd.com/pt/problems/view/1038
func main() {
	itensPrice:= map[int]float64 {
        1: 4.0,
        2: 4.5,
        3: 5.0,
        4: 2.0,
        5: 1.5,
    }
    
    var code, quantity int
    fmt.Scan(&code, &quantity)
    
    totalPrice := itensPrice[code] * float64(quantity)
    
    fmt.Printf("Total: R$ %.2f\n", totalPrice)
}