package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	Text    string
	Options []string
	Answer  string
}

func (q *Question) BuildQuestion(rawData []string) {
	q.Text = rawData[0]
	q.Options = rawData[1 : len(rawData)-1]
	q.Answer = rawData[len(rawData)-1]
}

type GameState struct {
	Name      string
	Score     int8
	Questions []Question
}

func (g *GameState) init() {
	fmt.Println("Seja bem vindo(a) ao quiz!")
	fmt.Println("Escreva seu nome: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name
	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProcessarCSV(done chan<- bool) {
	f, err := os.Open("quiz-go.csv")

	if err != nil {
		panic("Erro ao ler o arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao processar o arquivo")
	}

	for key, row := range records {
		if key == 0 {
			continue
		}

		question := &Question{}
		question.BuildQuestion(row)
		g.Questions = append(g.Questions, *question)
	}
	done <- true
}

func (g *GameState) Run() {
	answer := ""
	for key, question := range g.Questions {
		fmt.Printf("\n[%d] %s: \n", key+1, question.Text)
		for key, option := range question.Options {
			fmt.Printf("\t%d - %s\n", key+1, option)
		}

		fmt.Print("Resposta: ")
		fmt.Scan(&answer)

		if answer == question.Answer {
			g.Score += 10
			fmt.Println("Resposta Correta!")
		} else {
			fmt.Println("Resposta Incorreta!")
		}
	}
}

func (g *GameState) Finish() {
	correctAnswer := 0
	if g.Score > 0 {
		 correctAnswer = int(g.Score) / 10
	}
	fmt.Printf("\nVocê finalizou o quiz %s", g.Name)
	fmt.Printf("Respostas Corretas: %d | %d\n", int(correctAnswer), len(g.Questions))
	fmt.Printf("Pontuação Final: %d", g.Score);
}

func main() {
	game := &GameState{}
	cn := make(chan bool)
	go game.ProcessarCSV(cn)
	<-cn
	game.init()
	game.Run()
	game.Finish()
}
