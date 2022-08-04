package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// check if number of arguments = 1
	// if then, use RETURN func
	if len(os.Args) == 1 {
		fmt.Println("Give me a letter.")
		return
	}

	in := os.Args[1]
	// check if length of arguments = 1
	// use len(os.Args[1) to check
	if len(in) >= 2 {
		fmt.Println("Give me a letter, not a word.")
		return
	}

	//or use strings.IndexAny func in if statement directly
	// use strings.IndexAny to sort for vowels or semi-vowels in os.Args[1]
	if strings.IndexAny(in, "aeiou") == 0 {
		fmt.Printf("%q is a vowel.\n", in)
	} else if strings.IndexAny(in, "yw") == 0 {
		fmt.Printf("%q is a semi-vowel.\n", in)
	} else {
		fmt.Printf("%q is a consonant.\n", in)
	}

	//print results
}
