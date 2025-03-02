package main
 
import (
    "fmt"
)
 
func main() {
    
    var salario, reajuste float64
    var percentual int64

	salario = 400.00
    
    switch {
    case salario <= 400.00:
        percentual = 15
    case salario > 400.00 && salario <= 800.00:
        percentual = 12
    case salario > 800.00 && salario <= 1200.00:
         percentual = 10
    case salario > 1200.00 && salario <= 2000.00:
         percentual = 7
     case salario > 2000.00:
         percentual = 4
    }
    
    reajuste = salario * float64(percentual) / 100
    novoSalario := salario + reajuste
    fmt.Printf("Novo  salario: %.2f\n", novoSalario)
    fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
    fmt.Println("Em percentual:", percentual, "%")
    
}