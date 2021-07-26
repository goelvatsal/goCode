package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func main() {
	a := os.Args
	n, err := strconv.Atoi(a[1])

	if err != nil {
		fmt.Printf("That's not a age.\n")
	} else if len(a) != 2 {
		fmt.Printf("This program requires a age.\n")
		return
	} else if n > 17 {
		fmt.Printf("R-Rated\n")
	} else if n >= 13 && n <= 17 {
		fmt.Printf("PG-13\n")
	} else if n <= 12 {
		fmt.Printf("PG-Rated\n")
	} else if n < 0 {
		fmt.Printf("Wrong age: %d\n", n)
	}
}
