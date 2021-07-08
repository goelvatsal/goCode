package main

import ("fmt"; 	"path")

func main() {
	dir, _ := path.Split("css/main.css")

	fmt.Println("dir :", dir)
}