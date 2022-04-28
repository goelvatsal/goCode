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

const (
	header    = "Location,Size,Beds,Baths,Price"
	data      = "New York,100,2,1,100000,New York,150,3,2,200000,Paris,200,4,3,400000,Istanbul,500,10,5,1000000"
	separator = ","
)

func main() {
	dataS := strings.Split(data, separator)
	headerS := strings.Split(header, separator)
	printData(dataS, headerS)
	fmt.Printf("Average:       ")

	for i := 1; i < len(headerS); i++ {
		fmt.Printf("%-15.2f", findColumnAvg(dataS, headerS, i))
	}
	fmt.Println()
}

func findColumnData(dataS []string, headerS []string, desiredColumnIndex int) []string {
	var size []string

	var lines int
	for i := 0; i < len(dataS); i++ {
		if i == (desiredColumnIndex + lines*len(headerS)) {
			size = append(size, dataS[i])
		}

		if (i+1)%len(headerS) == 0 {
			lines++
		}
	}

	return size
}

func findAvg(nums []string) float64 {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		n, _ := strconv.ParseFloat(nums[i], 64)
		sum += n
	}

	//fmt.Println("abc", sum)
	return sum / float64(len(nums))
}

func findColumnAvg(dataS []string, headerS []string, desiredColumnIndex int) float64 {
	size := findColumnData(dataS, headerS, desiredColumnIndex)
	return findAvg(size)
}

func printData(dataS, headerS []string) {
	for i := 0; i < len(headerS); i++ {
		fmt.Printf("%-15s", headerS[i])
	}

	fmt.Println("\n===========================================================================")

	for i := 0; i < len(dataS); i++ {
		fmt.Printf("%-15s", dataS[i])

		if (i+1)%len(headerS) == 0 {
			fmt.Println()
		}
	}
	fmt.Println("\n===========================================================================")
}
