package main

import "fmt"

func main() {
	const (
		winter = 1
		summer = 3
	)

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Printf("books: %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("\nwinter: %#v\n", wBooks)
	fmt.Printf("\nsummer: %#v\n", sBooks)

	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}
}
