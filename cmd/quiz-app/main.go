package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse the optional flags
	filename := flag.String("csv", "problems.csv", "csv file")
	shuffle := flag.Bool("shuffle", false, "bool flag to shuffle csv file contents")
	flag.Parse()

	// Open csv file
	csvInfo := openCsv(*filename)

	// Run the quiz
	correctAnswers, exitEarly := quiz(csvInfo, *shuffle)
	if exitEarly {
		fmt.Println("User interrupted (Ctrl+C). Exiting!")
		os.Exit(1)
	}
	fmt.Println("Your final score is:", correctAnswers, "of", len(csvInfo))
}
