package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st Argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("3rd Argument:", os.Args[3])

	fmt.Println("Number of items inside of os.Args: ", len(os.Args))
}