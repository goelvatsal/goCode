package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	l := len(os.Args) - 1

	if l != 1 {
		fmt.Println("Usage: [feet]")
		return
	}

	n, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", os.Args[1])
		return
	}

	fmt.Printf("%g feet is %g meters.\n", n, n*.3048)
}
