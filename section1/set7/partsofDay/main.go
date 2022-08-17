package main

import (
	"fmt"
	"time"
)

func main() {
	switch t := time.Now().Hour(); {
	case t >= 18:
		fmt.Println("Good evening!")
	case t >= 14:
		fmt.Println("Good afternoon!")
	case t >= 6:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good night!")
	}
}
