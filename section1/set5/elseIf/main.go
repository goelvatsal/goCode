package main

import "fmt"

func main() {
	score := 4

	if score > 3 {
		fmt.Println("Good! :)")
	} else if score == 3 {
		fmt.Println("On the edge :|")
	} else if score == 2 {
		fmt.Println("Meh... :(")
	} else {
		fmt.Println("Low.")
	}
}
