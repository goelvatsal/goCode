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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func main() {
	// FOR THIS EXERCISE, LOOK AT PREVIOUS EXERCISE'S SOLUTION
	a := os.Args

	if len(a) != 2 {
		fmt.Println("Give me a year number.")
		return
	} else if year, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", a[1])
	} else if year%4 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
