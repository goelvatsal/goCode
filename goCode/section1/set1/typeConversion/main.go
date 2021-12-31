package main

import "fmt"

func main() {
	speed := 100 //speed is int
	force := 2.5 //force is float64

	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
