package main
 
import (
    "fmt"
)
 
func main() {

    var salario float32 = 3002.00
    
    var resultado float32
    switch {
    case salario <= 2000.00:
        fmt.Println("Isento");
    case salario > 2000.01 && salario <= 3000.00:
        resultado = salario * 0.8
    case salario > 3000.00 && salario <= 4500.00: 
         resultado = salario * 0.18
    case salario > 4500.00: 
        resultado = salario / 28 * 100
    }
    
    fmt.Printf("R$ %.2f\n", resultado)
    
}
