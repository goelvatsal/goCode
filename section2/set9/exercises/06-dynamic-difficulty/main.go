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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	var guess int
	var err error
	maxTurns := 5

	if args[0] == "-v" {
		guess, err = strconv.Atoi(args[1])
	} else {
		guess, err = strconv.Atoi(args[0])
	}

	switch {
	case guess > 10 && guess < 25:
		maxTurns = 8
	case guess >= 25 && guess < 50:
		maxTurns = 12
	case guess >= 50 && guess < 75:
		maxTurns = 20
	case guess >= 75:
		maxTurns = 30
	case guess <= 10:
		maxTurns = 5
	}

	if len(args) < 1 {
		fmt.Printf(`Welcome to the Lucky Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Want to play?`, maxTurns)
		fmt.Println()
		return
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
		for turn := 0; turn != maxTurns; turn++ {
			n := rand.Intn(guess + 1)

			fmt.Printf("%d ", n)

			if n == guess {
				fmt.Println("YOU WIN!")
				return
			}
		}
		fmt.Println("YOU LOST... Try again?")
		return
	} else {
		for turn := 0; turn != maxTurns; turn++ {
			n := rand.Intn(guess + 1)

			if n == guess {
				fmt.Println("YOU WIN!")
				return
			}
		}
		fmt.Println("YOU LOST... Try again?")
	}
}
