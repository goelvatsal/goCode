package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	const (
		maxTurns = 5
		usage    = `Welcome to the Lucky Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Want to play?`
	)

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		fmt.Println()
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[0])

	if err != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		n1 := rand.Intn(guess2 + 1)

		if turn == 0 && (n == guess || n1 == guess2) {
			fmt.Println("YOU WON ON YOUR FIRST TRY!")
			return
		}

		if n == guess || n1 == guess2 {
			fmt.Println("YOU WIN!")
			return
		}
	}
	fmt.Println("YOU LOST... Try again?")
}
