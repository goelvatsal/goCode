package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Provide an integer.")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}
}
