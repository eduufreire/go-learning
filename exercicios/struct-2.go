package main

import "fmt"

type Aluno struct {
	Nome      string
	Matricula int
	Nota      float64
}

func main() {

	alunos := []Aluno{
		Aluno{
			Nome:      "Edu",
			Matricula: 1,
			Nota:      9.9,
		},
		Aluno{
			Nome:      "Fulano",
			Matricula: 2,
			Nota:      4.8,
		},
		Aluno{
			Nome:      "Ciclano",
			Matricula: 3,
			Nota:      6.8,
		},
		Aluno{
			Nome:      "Beltrano",
			Matricula: 4,
			Nota:      8.1,
		},
		Aluno{
			Nome:      "Esquisito",
			Matricula: 5,
			Nota:      5.9,
		},
	}

	somatoriaNotas := 0.0
	for _, aluno := range alunos {
		somatoriaNotas += aluno.Nota
	}

	fmt.Printf("Média: %.1f", somatoriaNotas / float64(len(alunos)))

}