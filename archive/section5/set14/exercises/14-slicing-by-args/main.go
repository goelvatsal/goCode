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
// EXERCISE: Slicing by arguments
//
//   We've a []string, you will get arguments from the command-line,
//   then you will print only the elements that are requested.
//
//   1. Print the []string (it's in the code below)
//
//   2. Get the starting and stopping positions from the command-line
//
//   3. Print the []string slice by slicing it using the starting and stopping
//      positions
//
//   4. Handle the error cases (see the expected output below)
//
//   5. Add new elements to the []string slice literal.
//      Your program should work without changing the rest of the code.
//
//   6. Now, play with your program, get a deeper sense of how the slicing
//      works.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: we expect only two arguments)
//
//  go run main.go 1 2 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: starting index >= 0 && stopping pos <= len(slice) )
//
//  go run main.go -1 5
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  (error because: starting <= stopping)
//
//  go run main.go 3 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  go run main.go 0
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus Warsaw]
//
//
//  go run main.go 1
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus Warsaw]
//
//
//  go run main.go 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Nexus Warsaw]
//
//
//  go run main.go 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Warsaw]
//
//
//  go run main.go 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    []
//
//
//  go run main.go 0 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus]
//
//
//  go run main.go 1 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus]
//
//
//  go run main.go 1 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan]
//
// ---------------------------------------------------------

func main() {
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	//fmt.Println(len(os.Args))

	// print []string slice
	fmt.Printf("%q\n\n", ships)

	// handle errors
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Provide the [starting] and [stopping] positions.")
		//fmt.Println(os.Args)
		return
	}

	// get starting/stopping positions from command line
	start := os.Args[1]

	var endInt int
	if len(os.Args) == 3 {
		end := os.Args[2]
		endN, err1 := strconv.Atoi(end)

		if err1 != nil {
			fmt.Println("This program only accepts integer numbers.")
			return
		}

		if endN < 0 {
			fmt.Println("No negative numbers!")
			return
		}
		endInt = endN
	} else {
		endInt = len(ships)
	}

	startN, err := strconv.Atoi(start)

	if err != nil {
		fmt.Println("This program only accepts integer numbers.")
		return
	}

	if startN < 0 {
		fmt.Println("No negative numbers!")
		return
	}

	// use if statement or for loop to look using vars
	if os.Args[1] != "" {
		fmt.Printf("%q\n", ships[startN:endInt])
	} else {
		fmt.Printf("%q\n", ships[startN:])
	}
}
