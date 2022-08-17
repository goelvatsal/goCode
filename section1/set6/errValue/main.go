package main

import "strconv"
import "os"
import "fmt"

func main() {
	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Int value:  :", n)
	fmt.Println("Error Value :", err)
}
