// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	// set user as jack and pass as 1888
	user, pass := "jack", 1888

	// if no input received or only username or pass, print usage
	if len(os.Args) == 1 || len(os.Args) == 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	// if user and pass is incorrect print 'access denied for os.Args[1]'
	intVar, _ := strconv.Atoi(os.Args[2])
	if os.Args[1] != user && intVar != pass {
		fmt.Printf("Access denied for %q.\n", os.Args[1])
		// if user is correct but pass is wrong print 'invalid pass for os.Args[1]'
	} else if os.Args[1] == user && intVar != pass {
		fmt.Printf("Invalid password for %q.\n", os.Args[1])
		// if user and pass is correct print 'access granted to os.Args[1]'
	} else if os.Args[1] == user && intVar == pass {
		fmt.Printf("Access granted to %q.\n", os.Args[1])
	}
}
