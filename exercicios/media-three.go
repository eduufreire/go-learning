package main
 
import (
    "fmt"
)
 
func main() {

    var n1, n2, n3, n4 float64
    fmt.Scan(&n1, &n2, &n3, &n4)
    
    media := ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4 * 1)) / (2+3+4+1)
        
    fmt.Printf("Media: %.1f\n", media)
    if(media >= 5.0 && media <= 6.9) {
		fmt.Println("Aluno em exame.")
        exame(media)
    } else if media >= 7.0 {
        fmt.Println("Aluno aprovado.")
    } else {
        fmt.Println("Aluno reprovado.")
    }

}

func exame(media float64) {
    
    var notaExame float64
    fmt.Scan(&notaExame)
    
    mediaFinal := (media + notaExame) / 2
    
    fmt.Printf("Nota do exame: %.1f\n", notaExame)
    if(mediaFinal >= 5.0) {
         fmt.Println("Aluno aprovado.")
    } else {
        fmt.Println("Aluno reprovado.")
    }
    fmt.Printf("Media final: %.1f\n", mediaFinal)
    
    
}
