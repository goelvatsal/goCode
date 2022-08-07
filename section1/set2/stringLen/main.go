package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Vatsal"

	fmt.Println(utf8.RuneCountInString(name))
}
