// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
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

	if args[0] == "-v" {
		for turn := 0; turn < maxTurns; turn++ {
			var n int
			if guess < 10 {
				n = rand.Intn(11)
			} else {
				n = rand.Intn(guess + 1)
			}

			fmt.Printf("%d ", n)

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
	} else {
		for turn := 0; turn < maxTurns; turn++ {
			var n int
			if guess < 10 {
				n = rand.Intn(11)
			} else {
				n = rand.Intn(guess + 1)
			}

			if n == guess {
				fmt.Println("YOU WIN!")
				return
			}
		}
		fmt.Println("YOU LOST... Try again?")
	}
}
