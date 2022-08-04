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
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	var (
		names = [3]string{
			"Cory",
			"Sam",
			"Neel",
		}
		distances = [5]int{10, 5, 20, 15, 25}
		data      = [5]uint8{2, 10, 5, 8, 13}
		ratios    = [1]float64{76.52}
		alive     = [4]bool{true, true, true, true}
		zero      = [0]uint8{}
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
