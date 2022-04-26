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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza pizza
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

func main() {
	var (
		pizza     []string
		deptTimes []time.Time
		gradYears []int
		lights    []bool
	)

	pizza = append(pizza, "cheese", "bell peppers", "onions")
	deptTimes = append(deptTimes, time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*48), time.Now().Add(time.Hour*72))
	gradYears = append(gradYears, 10191, 2026, 1776)
	lights = append(lights, false, false, false)

	fmt.Printf("pizza: %s\n\n", pizza)
	fmt.Printf("gradYears: %d\n\n", gradYears)
	fmt.Printf("deptTimes: %s\n\n", deptTimes)
	fmt.Printf("lights: %t\n", lights)
}
