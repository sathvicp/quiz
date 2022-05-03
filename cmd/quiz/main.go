package main

import (
	"fmt"
	"sathvicp/quiz/internal/args"
	"sathvicp/quiz/internal/csv"
	"sathvicp/quiz/internal/structs"
)

func main() {
	var tLimit, csvFile = args.ParseArgs()
	fmt.Printf("%d\n", tLimit)

	score := 0

	problems := csv.GenerateProblems(csvFile)

	for _, problem := range problems {
		userAnswer := fetchUserAnswer(problem)
		if checkAnswer(problem, userAnswer) {
			score++
		}
	}

	fmt.Printf("\nYour score is %d out of %d\n", score, len(problems))
}

func fetchUserAnswer(prob structs.Problem) string {
	fmt.Printf("\nQ: %s\n", prob.Question)
	fmt.Printf("A: ")
	var userAnswer string
	fmt.Scanln(&userAnswer)
	return userAnswer
}

func checkAnswer(prob structs.Problem, answer string) bool {
	return answer == prob.Answer
}
