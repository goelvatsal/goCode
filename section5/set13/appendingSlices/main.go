package main

import "fmt"

func main() {
	var todo []string

	todo = append(todo, "sing")
	todo = append(todo, "run", "code", "play")
	todo = append(todo, "see mom", "learn go")

	fmt.Print("todo: ", todo, "\n")
}
