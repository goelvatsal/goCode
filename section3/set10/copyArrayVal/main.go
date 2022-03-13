package main

import "fmt"

func main() {
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Edition"
	}
	books[3] = "Awesomeness"

	fmt.Printf("last year's books: %#v\n", prev)
	fmt.Printf("\nthis year's books: %#v\n", books)
}
