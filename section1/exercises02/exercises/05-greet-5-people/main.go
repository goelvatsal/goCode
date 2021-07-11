package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	fmt.Println("There are", len(os.Args) - 1, "great people!")
	fmt.Println("Hello, Great", os.Args[1])
	fmt.Println("Hello, Great", os.Args[2])
	fmt.Println("Hello, Great", os.Args[3])
	fmt.Println("Hello, Great", os.Args[4])
	fmt.Println("Hello, Great", os.Args[5])
}
