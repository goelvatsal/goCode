package main

import (
	"fmt"
	"os"
)

func main() {
	m := os.Args[1] / 3.281

	fmt.Printf("%s feet is %s meters.\n", os.Args[1], m)
}