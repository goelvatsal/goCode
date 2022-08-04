package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	e := strings.Repeat("!", l)

	s := e + strings.ToUpper(msg) + e

	fmt.Println(s)
}
