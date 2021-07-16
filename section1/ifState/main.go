package main

import "fmt"

func main() {
	score := 2

	if score > 3 {
		fmt.Printf("good\n")
	} else if score == 3{
		fmt.Printf("on the edge\n")
	} else if score == 2 {
		fmt.Printf("meh...\n")
	} else {
		fmt.Printf("bad\n")
	}
}