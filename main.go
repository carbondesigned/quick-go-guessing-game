package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// produces a random number
	rand.Seed(time.Now().UnixNano())

	var MAX_NUMBER int = 100
	var MIN_NUMBER int = 1
	var guess int

	// generate a new random number between 1 and 100 each time the program runs
	var answer = rand.Intn(MAX_NUMBER - MIN_NUMBER)
	println("Guessing Game")

	fmt.Printf("Enter a number between %v and %v: ", MIN_NUMBER, MAX_NUMBER)
	fmt.Scanf("%v", &guess)

	calculateGuess(guess, answer)
	fmt.Println("You got it!")
}

// "If the guess is not equal to the answer, print a message and ask for another guess."
//
// The function takes two arguments: guess and answer. It returns a boolean value
func calculateGuess(guess int, answer int) bool {
	for guess != answer {
		if guess > answer {
			fmt.Printf("Too high. Try again: ")
		} else if guess < answer {
			fmt.Printf("Too low. Try again: ")
		}
		fmt.Scanf("%v", &guess)
	}
	return true
}
