package main

import (
	"fmt"
	"strconv"
	"time"
)

const questionTimeLimit = 10

type Question struct {
	QuestionText string
	Options      [4]string
	Answer       int
}

func askQuestion(q Question, questionNumber int) int {
	fmt.Printf("\nQuestion %d: %s\n", questionNumber, q.QuestionText)
	for i, option := range q.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	timer := time.NewTimer(time.Duration(questionTimeLimit) * time.Second)
	answerCh := make(chan int)

	go func() {
		var answer string
		fmt.Print("Enter your answer (1-4) or type 'exit' to quit: ")
		_, err := fmt.Scanln(&answer)
		if err != nil {
			answerCh <- -1
			return
		}
		if answer == "exit" {
			answerCh <- 999
			return
		}
		parsedAnswer, err := strconv.Atoi(answer)
		if err != nil || parsedAnswer < 1 || parsedAnswer > 4 {
			answerCh <- -1
			return
		}
		answerCh <- parsedAnswer - 1
	}()

	select {
	case <-timer.C:
		fmt.Println("\nTime is up!")
		return -1
	case answer := <-answerCh:
		if answer == 999 {
			return 999
		}
		if answer < 0 || answer > 3 {
			fmt.Println("Invalid answer, please try again.")
			return -1
		}
		return answer
	}
}

func calculateScore(questions []Question, answers []int) {
	correctAnswers := 0
	for i, q := range questions {
		if answers[i] == q.Answer {
			correctAnswers++
		}
	}

	fmt.Printf("\nYou answered %d out of %d questions correctly.\n", correctAnswers, len(questions))
	performance := ""
	switch {
	case correctAnswers == len(questions):
		performance = "Excellent"
	case correctAnswers >= len(questions)/2:
		performance = "Good"
	default:
		performance = "Needs Improvement"
	}
	fmt.Printf("Performance: %s\n", performance)
}

func takeQuiz(questions []Question) {
	var answers []int
	for i, q := range questions {
		answer := askQuestion(q, i+1)
		if answer == -1 {
			fmt.Println("Skipping to the next question due to invalid input.")
			continue
		}

		if answer == 999 {
			fmt.Println("Exiting the quiz early...")
			break
		}

		answers = append(answers, answer)
	}

	calculateScore(questions, answers)
}

func main() {
	questions := []Question{
		{
			QuestionText: "What is the capital of France?",
			Options:      [4]string{"Berlin", "Madrid", "Paris", "Rome"},
			Answer:       2,
		},
		{
			QuestionText: "Which programming language is known as 'Go'?",
			Options:      [4]string{"Java", "C", "Go", "Python"},
			Answer:       2,
		},
		{
			QuestionText: "What is 2 + 2?",
			Options:      [4]string{"3", "4", "5", "6"},
			Answer:       1,
		},
	}

	fmt.Println("Welcome to the Online Quiz System!")
	fmt.Println("Type 'exit' at any time to quit the quiz early.")

	takeQuiz(questions)
}
