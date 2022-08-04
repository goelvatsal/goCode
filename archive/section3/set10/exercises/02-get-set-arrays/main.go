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
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func main() {
	var (
		names     = [3]string{"Cory", "Sam", "Neel"}
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
