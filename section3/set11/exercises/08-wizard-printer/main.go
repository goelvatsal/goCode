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
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multidimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	scientists := [5][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	// debugging
	for i := 0; i < 5; i++ {
		fmt.Printf("scientists[%d]: %q\n", i, scientists[i])
	}

	fmt.Println("\nFirst Name    Last Name    Nickname")
	fmt.Println("===================================")

	fmt.Printf("%v        %v     %v\n", scientists[0][0], scientists[0][1], scientists[0][2])
	fmt.Printf("%v         %v       %v\n", scientists[1][0], scientists[1][1], scientists[1][2])
	fmt.Printf("%v       %v      %v\n", scientists[2][0], scientists[2][1], scientists[2][2])
	fmt.Printf("%v         %v        %v\n", scientists[3][0], scientists[3][1], scientists[3][2])
	fmt.Printf("%v       %v       %v\n", scientists[4][0], scientists[4][1], scientists[4][2])
}
