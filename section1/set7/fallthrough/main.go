package main

import "fmt"

func main() {
	i := 142

	switch {
	case i > 100:
		fmt.Print("Big ")
		fallthrough
	case i > 0:
		fmt.Print("Positive ")
		fallthrough
	default:
		fmt.Print("Number")
	}

	fmt.Println()
}
