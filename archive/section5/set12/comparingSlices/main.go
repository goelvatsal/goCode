package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "Dracula"
	books[1] = "1984"
	books[2] = "Island"

	newBooks := [5]string{"ulysses", "fire"}
	books = newBooks

	var games []string

	games = []string{"pokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}

	newGames = games
	games = nil

	var ok string
	for i, game := range games {
		for game != newGames[i] {
			ok = "not "
			break
		}
	}

	if len(games) != len(newGames) {
		ok = "not "
	}

	fmt.Printf("games and newGames are %sequal. \n\n", ok)
}
