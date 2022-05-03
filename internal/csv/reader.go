package csv

import (
	"encoding/csv"
	"log"
	"os"
	"sathvicp/quiz/internal/structs"
	"strings"
)

func getReader(filePath string) *csv.Reader {
	fileBuffer, ioErr := os.Open(filePath)

	if ioErr != nil {
		log.Fatal(ioErr)
	}
	return csv.NewReader(fileBuffer)
}

func GenerateProblems(filePath string) []structs.Problem {
	reader := getReader(filePath)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	problems := make([]structs.Problem, len(lines))

	for i, line := range lines {
		problems[i] = structs.Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}
