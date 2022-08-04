package main

import "fmt"

func main() {
	speed, force := 100, 2.5

	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
