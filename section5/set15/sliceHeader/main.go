package main

import (
	s "github.com/inancgumus/prettyslice"
	"unsafe"
)

func main() {
	type collection []string
	data := collection{"slices", "are", "awesome", "period", "!!"}
	change(data)

	s.Show("main's data:", data)
	s.Show("main's data slice address: %p\n", &data)

	array := [...]string{"slices", "are", "awesome", "period", "!!"}

	s.Show("array's size: %d bytes\n", unsafe.Sizeof(array))
	s.Show("data's size: %d bytes\n", unsafe.Sizeof(data))
}

func change(data []string) {
	data[2] = "brilliant!"
	s.Show("change's data:", data)
	s.Show("change's data slice address: %p\n", &data)

}
