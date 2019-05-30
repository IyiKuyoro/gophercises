package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/csv"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the SV file: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Could not parse the csv file")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			line[0],
			line[1],
		}
	}

	return problems
}

type problem struct {
	question string
	answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
