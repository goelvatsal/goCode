package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		fmt.Println()
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for t := 0; t < maxTurns; t++ {
		n := rand.Intn(guess + 1)

		if t == 1 && n == guess {
			fmt.Println("YOU WON ON YOUR FIRST TRY!")
			return
		}
		if n == guess {
			switch rand.Intn(2) {
			case 0:
				fmt.Println("YOU WIN!")
				return
			case 1:
				fmt.Println("PERFECT!")
				return
			}
		}
	}

	// this statement is to print random error messages
	switch rand.Intn(2) {
	case 0:
		fmt.Println("YOU LOST... Try again?")
		return
	case 1:
		fmt.Println("NICE TRY... Try again?")
		return
	}
}
