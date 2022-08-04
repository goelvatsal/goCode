package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])
	var uniques []int

loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}
	fmt.Println("\n\nUniques:", uniques)

	sort.Ints(uniques)
	fmt.Println("\nsorted:", uniques)
}
