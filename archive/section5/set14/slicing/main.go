package main

import "fmt"

func main() {
	items := []string{"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest", "ultima"}

	top3 := items[:3]
	fmt.Println("top 3:", top3)

	l := len(items)

	last4 := items[l-4:]
	fmt.Println("last4:", last4)

	mid := last4[1:3]
	fmt.Println("mid items:", mid)

	fmt.Println(items[9], last4[0])
}
