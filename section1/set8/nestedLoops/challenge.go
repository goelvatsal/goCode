package main

// input: an integer for what number to make the table for
// if input is 5; then program prints 5x5 table

// output: the table by which the user inputted the number
// input = 5; output is 5x5 table

// goal: print multiplication table

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Provide an integer.")
		return
	}

	//prints header column
	fmt.Printf("%5s", "X")
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// prints vertical column
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		// multiplication code
		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}
}
