package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	fmt.Printf("Converted number: %d\n", n)
	fmt.Printf("Returned error value: %s\n", err)
}