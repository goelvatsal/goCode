package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	l := len(args) - 1

	//	if nothing passed, then print usage
	if l == 0 {
		fmt.Printf("Usage: [username] [password]\n")
		return
	}

	//	user creds
	u := "jack"
	p := "1888"

	u2 := "inanc"
	p2 := "1879"

	if !(args[1] == u || args[1] == u2) {
		fmt.Printf("Access denied for %q.\n", args[1])
	}

	if args[1] == u && args[2] != p || args[1] == u2 && args[2] != p2 {
		fmt.Printf("Invalid password for user %q.\n", args[1])
	}

	if args[1] == u && args[2] == p || args[1] == u2 && args[2] == p2 {
		fmt.Printf("Access granted to user %q.\n", args[1])
	}
}
