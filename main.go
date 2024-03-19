package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	problems, err := readProblems(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to read problems from CSV: %v", err))
	}

	correct := conductQuiz(problems)
	fmt.Printf("You got %d out of %d correct\n", correct, len(problems))

}

func readProblems(filename string) ([]Problem, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		return nil, err
	}
	return parseLines(lines), nil
}

func conductQuiz(problems []Problem) int {
	correct := 0

	for i, p := range problems {
		fmt.Printf("#%d: %s = \n", i+1, p.Question)

		var answer string

		fmt.Scanf("%s\n", &answer)
		if answer == p.Answer {
			correct++
		}
	}

	return correct
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
