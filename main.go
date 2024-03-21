package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	// get and parse cli flags
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 60, "time limit (in seconds) to finish the quiz")

	flag.Parse()

	// generate problems
	problems, err := createProblems(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to read problems from CSV: %v", err))
	}

	// print a start msg
	startMsg(len(problems), *timeLimit)

	// start the quiz if the user writes 's'
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "s" {
		conductQuiz(problems, timeLimit)
	} else {
		exit("Exiting...")
	}
}

func conductQuiz(problems []Problem, timeLimit *int) int {
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("#%d: %s\n", i+1, p.Question)

		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			userInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			answerCh <- strings.ToLower(strings.TrimSpace(userInput))
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time's up! You got %d out of %d correct\n", correct, len(problems))
			return correct
		case answer := <-answerCh:
			if answer == p.Answer {
				correct++
			}
		}
	}

	fmt.Printf("You completed the quiz! You got %d out of %d correct\n", correct, len(problems))
	timer.Stop()

	return correct
}

func createProblems(filename string) ([]Problem, error) {
	parseLines := func(lines [][]string) []Problem {
		problems := make([]Problem, len(lines))

		for i, line := range lines {
			problems[i] = Problem{
				Question: line[0],
				Answer:   strings.TrimSpace(line[1]),
			}
		}

		return problems
	}

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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func startMsg(questionsAmount, timeLimit int) {
	fmt.Println("")
	fmt.Println("JavaScript Quiz")
	fmt.Println(`     ██╗███████╗     ██████╗ ██╗   ██╗██╗███████╗
     ██║██╔════╝    ██╔═══██╗██║   ██║██║╚══███╔╝
     ██║███████╗    ██║   ██║██║   ██║██║  ███╔╝
██   ██║╚════██║    ██║▄▄ ██║██║   ██║██║ ███╔╝
╚█████╔╝███████║    ╚██████╔╝╚██████╔╝██║███████╗
 ╚════╝ ╚══════╝     ╚══▀▀═╝  ╚═════╝ ╚═╝╚══════╝`)

	fmt.Printf("There are %d questions. There's a time limit of %d seconds. Write 's' and hit 'Enter' to start the quiz.", questionsAmount, timeLimit)
	fmt.Println("")
}
