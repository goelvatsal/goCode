package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted value    :", n)
	fmt.Println("Returned err value :", err)
}
