package main

import (
	"fmt"
	"os"
)

func main() {

	switch m := os.Args[1]; m {
	case "December", "January", "February":
		fmt.Println("Winter!")
	case "March", "April", "May":
		fmt.Println("Spring!")
	case "June", "July", "August":
		fmt.Println("Summer!")
	case "September", "October", "November":
		fmt.Println("Fall!")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}
