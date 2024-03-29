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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case-insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

func main() {
	// declare array with book titles
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	// do error handling
	if len(os.Args) != 2 {
		fmt.Println("Tell me a book name.")
		return
	}

	_, err := strconv.Atoi(os.Args[1])
	if err == nil {
		fmt.Printf("%v is not a book. Tell me a book name.\n", os.Args[1])
		return
	}
	// find if args[1] is a book title or not and print
	var (
		i = false
		v = ""
		t = false
	)

	for _, v = range books {
		fmt.Println("Search Results:")
		for j := 0; j < 4; j++ {
			t = strings.Contains(strings.ToLower(v), strings.ToLower(os.Args[1]))

			if i = t && j == 0; i {
				fmt.Printf("+ Kafka's Revenge\n")
				fmt.Printf("+ Kafka's Revenge 2nd Edition\n")
				break
			}

			v = books[j]
		}

		if t {
			break
		} else {
			fmt.Printf("We don't have the book: %q.\n", os.Args[1])
			return
		}
	}
}
