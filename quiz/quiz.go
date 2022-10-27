package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func quiz(records [][]string, score *int, c chan bool) {
	var ans string

	for _, row := range records {
		fmt.Println(row[0])
		fmt.Scanln(&ans)
		if ans == row[1] {
			*score++
		}
	}
	close(c)
}

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file.")
	}

	score := 0
	done := make(chan bool)

	fmt.Println("Hit Enter to start the quiz")
	fmt.Scanln()

	go quiz(records, &score, done)

	select {
	case <-done:
		break
	case <-time.After(time.Duration(time.Second * 30)):
		fmt.Println("\nTimes up!")
	}

	fmt.Printf("Score: %v out of %v\n", score, len(records))
}
