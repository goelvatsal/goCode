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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = "New York,100,2,1,100000,New York,150,3,2,200000,Paris,200,4,3,400000,Istanbul,500,10,5,1000000"

		separator = ","
	)

	headerS := strings.Split(header, separator)
	for i := 0; i < len(headerS); i++ {
		fmt.Printf("%-15s", headerS[i])
	}
	fmt.Println("\n===========================================================================")

	dataS := strings.Split(data, separator)
	for i := 0; i < len(dataS); i++ {
		fmt.Printf("%-15s", dataS[i])

		if (i+1)%len(headerS) == 0 {
			fmt.Println()
		}
	}

	fmt.Println("\n===========================================================================")
	fmt.Printf("Average:       ")
	size := []string{dataS[1], dataS[6], dataS[11], dataS[16]}
	var sum float64
	for i := 0; i < len(size); i++ {
		n, _ := strconv.ParseFloat(size[i], 64)
		sum += n
	}
	avg := sum / float64(len(size))
	fmt.Printf("%-15.2f", avg)

	bed := []string{dataS[2], dataS[7], dataS[12], dataS[17]}
	sum = 0
	for i := 0; i < len(bed); i++ {
		n, _ := strconv.ParseFloat(bed[i], 64)
		sum += n
	}
	avg = sum / float64(len(bed))
	fmt.Printf("%-15.2f", avg)

	bath := []string{dataS[3], dataS[8], dataS[13], dataS[18]}
	sum = 0
	for i := 0; i < len(bath); i++ {
		n, _ := strconv.ParseFloat(bed[i], 64)
		sum += n
	}
	avg = sum / float64(len(bath))
	fmt.Printf("%-15.2f", avg)

	price := []string{dataS[4], dataS[9], dataS[14], dataS[19]}
	sum = 0
	for i := 0; i < len(price); i++ {
		n, _ := strconv.ParseFloat(price[i], 64)
		sum += n
	}
	avg = sum / float64(len(price))
	fmt.Printf("%-15.0f\n", avg)
}
