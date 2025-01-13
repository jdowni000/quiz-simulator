package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

// quiz allows 30 seconds to answer all questions provided
func quiz(problemInfo [][]string) (int, bool) {
	score := 0
	for _, problem := range problemInfo {
		// Create a channel to receive OS signals (e.g., Ctrl+C)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		// Start the quiz timer
		timer := time.NewTimer(30 * time.Second)

		fmt.Println("What is", problem[0], "?")
		// Create a channel to receive the answer
		answerCh := make(chan string)

		// Get user input in a separate goroutine
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			userAnswer = strings.TrimSpace(userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case answer := <-answerCh:
			// User provided an answer
			if answer == problem[1] {
				score++
				fmt.Println("*** That is CORRECT! ***")
			} else {
				fmt.Println("*** That is INCORRECT! ***")
			}
		case <-timer.C:
			// Timer expired
			fmt.Println("Time's up!")
			return score, false
		case <-c:
			// User interrupted (Ctrl+C)
			return score, true
		}
	}
	return score, false
}
