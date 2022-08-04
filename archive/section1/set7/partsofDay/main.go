package main

import (
	"fmt"
	"time"
)

func main() {

	switch t := time.Now().Hour(); {
	case t > 9 && t < 12:
		fmt.Println("Good morning!")
	case t >= 12 && t <= 18:
		fmt.Println("Good afternoon!")
	case t > 18 && t <= 20:
		fmt.Println("Good evening!")
	case t >= 0 && t <= 9:
		fmt.Println("Good night!")
	}
}
