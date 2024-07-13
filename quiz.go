package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Question struct represents a single quiz question
type Question struct {
	Question string
	Answer   int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var correctAnswers int
	numQuestions := 5
	maxAttempts := 3

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Attempt %d:\n", attempt)

		// Generate new set of questions for each attempt
		questions := generateQuestions(numQuestions)

		correctAnswers = conductQuiz(questions)

		fmt.Printf("\nAttempt %d ended. You answered %d out of %d questions correctly.\n\n", attempt, correctAnswers, numQuestions)

		if attempt < maxAttempts {
			var response string
			fmt.Printf("Do you want to attempt the quiz again? (yes/no): ")
			fmt.Scanln(&response)

			if response != "yes" && response != "y" {
				break
			}
		}
	}

	fmt.Println("Thank you for playing!")
}

// generateQuestions creates and shuffles a new set of quiz questions
func generateQuestions(numQuestions int) []Question {
	questions := []Question{
		{"What is 5 + 3?", 8},
		{"What is 10 - 6?", 4},
		{"What is 7 * 2?", 14},
		{"What is 15 / 3?", 5},
		{"What is 12 + 9?", 21},
	}

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	return questions[:numQuestions]
}

// conductQuiz conducts the quiz based on the provided questions and returns the number of correct answers
func conductQuiz(questions []Question) int {
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()

	correctAnswers := 0

questionLoop:
	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question.Question)

		answerCh := make(chan int)
		go func() {
			var answer int
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			break questionLoop
		case answer := <-answerCh:
			if answer == question.Answer {
				fmt.Println("Correct!\n")
				correctAnswers++
			} else {
				fmt.Printf("Wrong! The correct answer is %d\n\n", question.Answer)
			}
		}
	}

	return correctAnswers
}
