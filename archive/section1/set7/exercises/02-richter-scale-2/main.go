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
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go aliens
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

func main() {
	// handle error checking
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake in human terms.")
		return
	}

	//
	//// handle error checking
	//if err != nil {
	//	fmt.Printf("%s's richter scale is unknown.\n", os.Args[1])
	//	return
	//}

	// create switch statement
	switch {
	case os.Args[1] == "micro":
		fmt.Printf("micro's richter scale is less than 2.0\n")
	case os.Args[1] == "very minor":
		fmt.Printf("very minor's richter scale is 2 - 2.9\n")
	case os.Args[1] == "minor":
		fmt.Printf("minor's richter scale is 3 - 3.9\n")
	case os.Args[1] == "light":
		fmt.Printf("light's richter scale is 4 - 4.9\n")
	case os.Args[1] == "moderate":
		fmt.Printf("moderate's richter scale is 5 - 5.9\n")
	case os.Args[1] == "strong":
		fmt.Printf("strong's richter scale is 6 - 6.9\n")
	case os.Args[1] == "major":
		fmt.Printf("major's richter scale is 7 - 7.9\n")
	case os.Args[1] == "great":
		fmt.Printf("great's richter scale is 8 - 9.9\n")
	case os.Args[1] == "massive":
		fmt.Printf("massive's richter scale is 10+\n")
	}
}
