package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	// if no user input received, then print usage msg
	if len(os.Args) != 2 || args[0] == "" {
		fmt.Println("[your name]")
		return
	}

	// error handling (ex: user types in int)
	n, err := strconv.Atoi(args[0])
	if err == nil {
		fmt.Printf("%d is not a name.\n", n)
	}

	// user time in nanoseconds as seed
	rand.Seed(time.Now().UnixNano())

	// type emotions in array
	emotions := [...]string{
		"sad",
		"terrible",
		"happy",
		"awesome",
		"good",
	}

	// use switch statement to print emotion + name
	fmt.Printf("%s feels %s!\n", args[0], emotions[rand.Intn(len(emotions))])
}
