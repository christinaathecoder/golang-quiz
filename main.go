package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Problem struct w/ question && answer
type Problem struct {
	q string
	a string
}

func main() {
	//flag package
	fileName := flag.String("csv", "problems.csv", "CSV file formatted as 'question,answer'")
	flag.Parse()

	//fileName is pointer to string, use * to get value from string
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *fileName))
	}

	//read in csv file
	r := csv.NewReader(file)
	//slice
	rows, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse CSV file\n")
	}
	//get slice of problems
	problems := parseRows(rows)
	//counter for num of correct answers
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q) //i+1 to start at 1

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.a {
			correct++
		}
	}
	fmt.Printf("You got %d/%d correct!\n", correct, len(problems))
}

// parse rows - create problem with question and answer
func parseRows(rows [][]string) []Problem {
	ret := make([]Problem, len(rows))
	for i, row := range rows {
		ret[i] = Problem{
			q: row[0],
			a: row[1],
		}
	}
	return ret
}

// exit message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
