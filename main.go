package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem struct w/ question && answer
type Problem struct {
	q string
	a string
}

func main() {
	//flag package
	fileName := flag.String("csv", "problems.csv", "CSV file formatted as 'question,answer'")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz")
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
	//timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//counter for num of correct answers
	correct := 0

	for i, problem := range problems {
		//print problem
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q) //i+1 to start at 1
		answerChannel := make(chan string)
		//go routine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			//send answer to answer channel
			answerChannel <- answer
		}()
		select {
		case <-timer.C: //if we get answer from timer channel
			fmt.Printf("\nYou got %d/%d correct!\n", correct, len(problems))
			return
		case answer := <-answerChannel: //if we get answer from answer channel
			if answer == problem.a { //check if answer is correct
				correct++
			}
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
			a: strings.TrimSpace(row[1]),
		}
	}
	return ret
}

// exit message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
