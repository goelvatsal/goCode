// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	var (
		names     = [...]string{"Cory", "Sam", "Neel"}
		distances = [...]int{10, 5, 20, 15, 25}
		data      = [...]uint8{2, 10, 5, 8, 13}
		ratios    = [...]float64{76.52}
		alive     = [...]bool{true, true, true, true}
		zero      = [...]uint8{}
	)

	for i := 0; i < 3; i++ {
		fmt.Printf("names[%d]: %v\n", i, names[i])
	}

	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("distances[%d]: %v\n", i, distances[i])
	}

	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("data[%d]: %v\n", i, data[i])
	}

	fmt.Println()
	fmt.Printf("ratios[0]: %v\n", ratios[0])

	fmt.Println()
	for i := 0; i < 4; i++ {
		fmt.Printf("alives[%d]: %v\n", i, alive[i])
	}

	fmt.Println()
	fmt.Printf("zero[0]: %v\n\n", zero)

	for i, q := range names {
		fmt.Printf("names[%d]: %s\n", i, q)
	}

	fmt.Println()
	for i, q := range distances {
		fmt.Printf("distances[%d]: %v\n", i, q)
	}

	fmt.Println()
	for i, q := range data {
		fmt.Printf("data[%d]: %v\n", i, q)
	}

	fmt.Println()
	for i, q := range ratios {
		fmt.Printf("ratios[%d]: %v\n", i, q)
	}

	fmt.Println()
	for i, q := range alive {
		fmt.Printf("alive[%d]: %v\n", i, q)
	}

	fmt.Println()
	fmt.Printf("zero[0]: %v\n", zero)
}
