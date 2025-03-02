package main
 
import (
    "fmt"
    "math"
)
 
func main() {

	var auxiliar1, auxiliar2, auxiliar3 = 7.0, 7.0, 5.0
    var a, b, c float64 
    
	    
    naoTriangulo := (a >= b + c)
    trianguloRetangulo := (math.Pow(a, 2) == math.Pow(b, 2) + math.Pow(c, 2))
    trianguloObtusangulo := (math.Pow(a, 2) > math.Pow(b, 2) + math.Pow(c, 2))
    trianguloAcutangulo := (math.Pow(a, 2) < math.Pow(b, 2) + math.Pow(c, 2))
    trianguloEquilatero := (a == b && b == c)
    trianguloIsosceles := ((a == b &&  b != c) || (a == c && c != b) || (b == c && b != a))

    if naoTriangulo {
        fmt.Println("NAO FORMA TRIANGULO")
    }
    
    if trianguloRetangulo {
        fmt.Println("TRIANGULO RETANGULO")
    }
    
    if trianguloObtusangulo {
        fmt.Println("TRIANGULO OBTUSANGULO")
    }
    
    if trianguloAcutangulo {
        fmt.Println("TRIANGULO ACUTANGULO")
    }
    
    if trianguloEquilatero {
        fmt.Println("TRIANGULO EQUILATERO")
    }
    
    if trianguloIsosceles {
        fmt.Println("TRIANGULO ISOSCELES")
    }
    
}