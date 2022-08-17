package main

import "os"
import "fmt"
import "strconv"

func main() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}
