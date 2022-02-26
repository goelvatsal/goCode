package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
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
	var guess int
	var err error

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		fmt.Println()
		return
	}

	if args[0] == "-v" {
		guess, err = strconv.Atoi(args[1])
	} else {
		guess, err = strconv.Atoi(args[0])
	}
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if args[0] == "-v" {
			fmt.Printf("%d ", n)
		}

		if turn == 4 || n == guess {
			fmt.Println()
		}

		if n == guess {
			fmt.Println("YOU WIN!")
			return
		}
	}
	fmt.Println("YOU LOST... Try again?")
	return
}
