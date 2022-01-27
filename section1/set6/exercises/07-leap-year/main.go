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
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	// declare variables
	a := os.Args

	// check if os.Args[1] is empty
	if len(a) != 2 {
		fmt.Println("Give me a year number.")
		return
	}

	// check if no number input is provided
	if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", a[1])
		// check if number is leap year if divisible by 4 cleanly
	} else if n%4 == 0 {
		fmt.Printf("%d is a leap year.\n", n)
		// check if number is leap year if divisible by 4 cleanly
	} else if n%4 != 0 {
		fmt.Printf("%d is not a leap year.\n", n)
	}

}
