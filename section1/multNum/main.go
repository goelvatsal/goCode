package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := os.Args
	if len(a) != 2 {
		fmt.Printf("Give me a number.\n")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		fmt.Printf("%s * 2 = %d\n", a[1], n*2)
	}
}