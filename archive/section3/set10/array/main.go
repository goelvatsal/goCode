package main

import "fmt"

func main() {
	const (
		winter = 1
		summer = 3
		yearly = winter + summer
	)

	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] = books[0] + " 2nd Edition"

	fmt.Printf("books: %#v\n", books)
}
