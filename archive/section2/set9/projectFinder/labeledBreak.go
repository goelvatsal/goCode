package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const corpus = "" + "lazy cat jumps again and again and again"
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break queries
			}
		}
	}
}
