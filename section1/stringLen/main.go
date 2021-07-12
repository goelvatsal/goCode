package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	s := os.Args[1]
	fmt.Println(utf8.RuneCountInString(s))
}