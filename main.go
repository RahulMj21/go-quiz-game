package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("file", "problems.csv", "File contains all the problems")
	timer := flag.Int("time", 30, "Time to solve the Problems")

	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("unable to open the file : %s", *fileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("unable to parse the file : %s", *fileName))
	}

	problems := parseLines(lines)
	timeLimit := time.NewTimer(time.Duration(*timer) * time.Second)
	correct := 0

	fmt.Println("RadheShyam ❣️, Let's start the quiz 🎉")
	for i, problem := range problems {
		fmt.Printf("Problem No #%d: %s = ", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timeLimit.C:
			{
				fmt.Printf("\nYou ran out of time, your score is %d out of %d\n", correct, len(problems))
				return
			}
		case answer := <-answerCh:
			{
				if answer == problem.answer {
					correct++
				}

			}
		}
	}
	fmt.Printf("You have scored %d out of %d\n", correct, len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i].question = line[0]
		problems[i].answer = strings.TrimSpace(line[1])
	}

	return problems
}
