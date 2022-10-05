package problem

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCsv(file string) []Problem {
	//Reading from csv file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true

	var problems []Problem

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(line)
		problems = append(problems, Problem{line[0], line[1]})
	}

	// fmt.Println(problems)
	// for _, p := range problems {
	// 	fmt.Println("Problem:", p.Question, "Answer:", p.Answer)
	// }

	return problems
}
