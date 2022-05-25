package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var MAX_NUMBER int = 100
	var MIN_NUMBER int = 1
	var guess int

	// generate a new random number between 1 and 100 each time the program runs
	var answer = rand.Intn(MAX_NUMBER-MIN_NUMBER+1) + MIN_NUMBER
	println(answer)
	println("Guessing Game")

	fmt.Printf("Enter a number between %v and %v: ", MIN_NUMBER, MAX_NUMBER)
	fmt.Scanf("%v", &guess)

	calculateGuess(guess, answer)
	fmt.Println("You got it!")
}

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
