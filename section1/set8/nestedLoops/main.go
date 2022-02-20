package main

import "fmt"

func main() {
	const max = 5

<<<<<<< HEAD
	// first header column; .g. x 0 1 2 3
=======
	// first header column; e.g. x 0 1 2 3
>>>>>>> 43d0476d93aa29c12b832d765c93c0f6360f7048
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// vertical column
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		// multiplication code; horizonal column
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}
}
