package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	m := float64(n)
	fmt.Printf("%.2f feet is %.2f meters.\n", m, m/3.281)
}
