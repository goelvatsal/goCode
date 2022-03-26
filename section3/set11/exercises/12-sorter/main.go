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
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func main() {
	// declare array
	a := [...]int{0, 0, 0, 0, 0}

	if len(os.Args) != 6 {
		fmt.Println("Please give me 5 numbers.")
		return
	}

	// error handling
	n, err := strconv.Atoi(os.Args[1])
	n1, err1 := strconv.Atoi(os.Args[2])
	n2, err2 := strconv.Atoi(os.Args[3])
	n3, err3 := strconv.Atoi(os.Args[4])
	n4, err4 := strconv.Atoi(os.Args[5])

	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Please give me 5 numbers.")
		return
	}

	arrN := [5]int{n, n1, n2, n3, n4}

	for i := 1; i < 6; i++ {
		a[i] = arrN[i]
	}
	fmt.Printf("%v\n", a)

	// add bubble sorting program
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	// print the sorted array
}
