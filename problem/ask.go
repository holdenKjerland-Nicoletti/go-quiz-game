package problem

import (
	"fmt"
)

func (p Problem) Ask() bool {
	var guess string

	fmt.Println(p.Question)

	_, err := fmt.Scanln(&guess)
	if err != nil || guess != p.Answer {
		return false
	}
	return true
}

func AskAll(problems []Problem, correct *int, done chan<- bool) {

	for _, prob := range problems {
		if prob.Ask() {
			*correct++
		}
	}

	done <- true
}
