package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type Problem struct {
	q string
	a string
}

func parseProblems(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, p := range lines {
		problems[i] = Problem{p[0], strings.TrimSpace(p[1])}
	}
	return problems
}

func main() {
	csvFileName := flag.String("problem", "problems.csv", "Input file containing question answers in format question,answer")
	flag.Parse()

	csvFile, err := os.Open(path.Join("Excercise1", *csvFileName))
	if err != nil {
		exit(fmt.Sprintf("failed to open provided csv file %s: %s", *csvFileName, err))
	}

	csvReader := csv.NewReader(csvFile)
	lines, err := csvReader.ReadAll()
	if err != nil {
		exit("failed to parse provided csv file %s: %s")
	}
	problems := parseProblems(lines)
	var correct = 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var ans string
		_, err := fmt.Scanf("%s\n", &ans)
		if err != nil {
			exit(fmt.Sprintf("failed to read input: %s", err))
		}
		if ans == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
