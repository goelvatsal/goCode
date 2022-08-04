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
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	if u != "jack" && u != "inanc" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if u == "jack" && p == "1888" {
		fmt.Printf("Access granted to %q.\n", u)
	} else if u == "inanc" && p == "1879" {
		fmt.Printf("Access granted to %q.\n", u)
	} else {
		fmt.Printf("Invalid password for %q.\n", u)
	}

	// SWITCH STARTS HERE
	switch {
	case u != "jack" && u != "inanc":
		fmt.Printf("Access denied for %q.\n", u)
	case u == "jack" && p == "1888":
		fallthrough
	case u == "inanc" && p == "1879":
		fmt.Printf("Access granted for %q.\n", u)
	default:
		fmt.Printf("Invalid password for %q.\n", u)
	}
}
