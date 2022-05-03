package main

import (
	"fmt"
	"sathvicp/quiz/internal/args"
	"sathvicp/quiz/internal/csv"
	"sathvicp/quiz/internal/structs"
	"time"
)

func main() {
	var tLimit, csvFile = args.ParseArgs()

	score := 0

	problems := csv.GenerateProblems(csvFile)
	ansChan := make(chan string)
	var userAnswer string

	timer := time.NewTimer(time.Duration(tLimit) * time.Second)

	for _, problem := range problems {
		go popQAndGetAns(problem, ansChan)
		select {
		case <-timer.C:
			fmt.Printf("\nTime's up! You scored %d out of %d\n", score, len(problems))
			return
		case userAnswer = <-ansChan:
			checkAnswerAndUpdateScore(problem, userAnswer, &score)
		}
	}

	fmt.Printf("\nYou scored %d out of %d\n", score, len(problems))
}

/*
Pops a question, fetches the answer and pushes into answer channel
*/
func popQAndGetAns(prob structs.Problem, ansChan chan string) {
	fmt.Printf("Q: %s\nA: ", prob.Question)
	var userAnswer string
	fmt.Scanln(&userAnswer)
	fmt.Println()
	ansChan <- userAnswer
}

func checkAnswerAndUpdateScore(prob structs.Problem, answer string, score *int) {
	if answer == prob.Answer {
		*score++
	}
}
