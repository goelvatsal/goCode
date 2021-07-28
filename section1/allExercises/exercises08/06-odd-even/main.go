package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number.\n", os.Args[1])
		return
	}

	if n%8 == 0 {
		fmt.Printf("%d is a even number and it's divisible by 8.\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number.\n", n)
	} else {
		fmt.Printf("%d is a odd number.\n", n)
	}

}
