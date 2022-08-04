package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields("lazy cat jumps again and again and again")

	var (
		i int
		v string
	)

	for i, v = range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}

	fmt.Printf("Last value of i = %d, and v = %q.\n", i, v)
}
