package main

import (
	"fmt"
	"os"
)

func main() {
	csvInfo := openCsv()
	correctAnswers, exitEarly := quiz(csvInfo)
	if exitEarly {
		fmt.Println("User interrupted (Ctrl+C). Exiting!")
		os.Exit(1)
	}
	fmt.Println("Your final score is:", correctAnswers, "of", len(csvInfo))
}
