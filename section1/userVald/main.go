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
	}

	//	user cred
	username := "jack"
	password := "1888"

	if args[1] != username {
		fmt.Printf("Access denied for %q.\n", args[1])
	}

	if args[1] == username && args[2] != password {
		fmt.Printf("Invalid password for user %q.\n", args[1])
	}

	if args[1] == username && args[2] == password {
		fmt.Printf("Access granted to user %q.\n", args[1])
	}
}
