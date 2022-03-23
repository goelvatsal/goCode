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
	if len(os.Args) != 3 || args[0] == "" {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	// error handling (ex: user types in int)
	n, err := strconv.Atoi(args[0])
	if err == nil {
		fmt.Printf("%d is not a name.\n", n)
		return
	}

	// user time in nanoseconds as seed
	rand.Seed(time.Now().UnixNano())

	// type emotions in array
	emotions := [...][3]string{
		{"sad", "terrible", "bad"},
		{"happy", "awesome", "good"},
	}

	// use if statement to print emotion + name
	if args[1] == "positive" {
		//fmt.Printf("%s feels %s!\n", args[0], emotions[rand.Intn(len(emotions[0]))])
		fmt.Printf("%s feels %s!\n", args[0], emotions[1][rand.Intn(len(emotions[1]))])
	} else if args[1] == "negative" {
		fmt.Printf("%s feels %s!\n", args[0], emotions[0][rand.Intn(len(emotions[1]))])
	}
}
