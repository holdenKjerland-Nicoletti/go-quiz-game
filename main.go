package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/holdenKjerland-Nicoletti/go-quiz-game/problem"
)

func main() {

	// Getting input file
	inputPtr := flag.String("input", "problems.csv", "Input CSV file with problems")
	timeSecs := flag.Int("timer", 30, "Number of seconds for quiz")
	flag.Parse()

	fmt.Println("Reading from file: ", *inputPtr)
	fmt.Println("You will have", *timeSecs, "seconds")

	problems := problem.ReadCsv(*inputPtr)
	correct := 0

	fmt.Println("Press ENTER to start quiz")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*timeSecs) * time.Second)
	finished := make(chan bool)

	go problem.AskAll(problems, &correct, finished)

	select {
	case <-timer.C:
	case <-finished:
	}
	fmt.Println("Correct Answers:", correct, "Total Questions:", len(problems))
}
