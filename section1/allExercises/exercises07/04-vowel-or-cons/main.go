package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {
	args := os.Args
	l := len(args) - 1

	if l != 1 {
		fmt.Printf("Give me a letter.\n")
	} else if args[1] == "a" || args[1] == "e" || args[1] == "i" || args[1] == "o" || args[1] == "u" {
		fmt.Printf("%q is a vowel.\n", args[1])
	} else if args[1] == "y" || args[1] == "w" {
		fmt.Printf("%q sometimes is a vowel, sometimes not.\n", args[1])
	} else {
		fmt.Printf("%q is a consonant.\n", args[1])
	}
}
