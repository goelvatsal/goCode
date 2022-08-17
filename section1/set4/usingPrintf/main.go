package main

import "fmt"

func main() {
	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v million kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	fmt.Printf("%v is %v million kms away. Think! %[2]v million kms! %[1]v OMG!\n", planet, distance)
}
