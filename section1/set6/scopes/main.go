package main

import "os"
import "fmt"
import "strconv"

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.ParseFloat(a[1], 64); err != nil {
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%g * 2 = %g.\n", n, n*2)
	}
}
