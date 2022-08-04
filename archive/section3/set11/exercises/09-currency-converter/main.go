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
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the constants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go
//   invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     %.2f USD is 0.88 EUR
//     %.2f USD is 0.78 GBP
//     %.2f USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		EUR = 9 - iota
		GBP
		JPY
	)

	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	// declare variables
	args := os.Args[1:]

	if len(os.Args) != 2 || args[0] == "" {
		fmt.Println("Type in a number.")
		return
	}

	// convert os.Args[1] into int
	n, err := strconv.ParseFloat(args[0], 64)

	// error handling
	if err != nil {
		fmt.Printf("%v is an invalid amount. It should be a number.\n", args[0])
		return
	}

	// multiply float with currency ratio and print
	fmt.Printf("%.1f USD is %.2f EUR.\n", n, n*rates[EUR])
	fmt.Printf("%.1f USD is %.2f GBP.\n", n, n*rates[GBP])
	fmt.Printf("%.1f USD is %.2f JPY.\n", n, n*rates[JPY])
}
