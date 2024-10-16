package main

import "fmt"

func main() {

	var randomNumber int = 8
	var userGuess int

	var gameOver bool = false
	var userWin bool = false
	var tryCount int = 3
	fmt.Printf("Welcome to the number guessing game!\n\n")
	fmt.Printf("Game started! ")

	for gameOver == false && tryCount > 0 {
		fmt.Printf("Enter your guess: ")
		fmt.Scan(&userGuess)

		if userGuess == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number.\n")
			gameOver = true
			userWin = true
		} else {
			fmt.Printf("Try again!\n")
			tryCount--
			userWin = false
		}
	}
	if userWin {
		fmt.Printf("Game over, you win!\n")
	} else {
		fmt.Printf("Game over, you lost!\n")
	}

}
