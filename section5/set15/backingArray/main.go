package main

import (
	s "github.com/inancgumus/prettyslice"
	"sort"
)

func main() {
	s.MaxPerLine = 6

	grades := []float64{40, 10, 20, 50, 60, 70}
	newGrades := append([]float64(nil), grades...)

	front := newGrades[:3]
	front2 := front[:3]
	front3 := front

	sort.Float64s(front)

	s.Show("grades", grades[:])
	s.Show("newGrades", newGrades)
	s.Show("front", front)
	s.Show("front2", front2)
	s.Show("front3", front3)

}
