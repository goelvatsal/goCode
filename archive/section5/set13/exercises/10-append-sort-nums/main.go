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
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Append and Sort Numbers
//
//  We'll have a []string that should contain numbers.
//
//  Your task is to convert the []string to an int slice.
//
//  1. Get the numbers from the command-line
//
//  2. Append them to an []int
//
//  3. Sort the numbers
//
//  4. Print them
//
//  5. Handle the error cases
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    provide a few numbers
//
//  go run main.go 4 6 1 3 0 9 2
//    [0 1 2 3 4 6 9]
//
//  go run main.go a b c
//    []
//
//  go run main.go 4 a 1 c d 9
//    [1 4 9]
//
// ---------------------------------------------------------

func main() {
	// handle error handling
	if len(os.Args) == 1 {
		fmt.Println("Please provide some numbers.")
		return
	}

	// attach user inputted numbers to an int slice
	// because os.Args[1] is the first user-inputted number,
	// this program is starting from that
	var nums []string
	for i := 1; i < len(os.Args); i++ {
		nums = append(nums, os.Args[i])
	}

	// sort numbers using sort.Ints
	sort.Strings(nums)
	fmt.Println(nums)
}
