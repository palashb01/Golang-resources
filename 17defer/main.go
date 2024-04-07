package main

import "fmt"

func main() {
	// Hello, World 2 will print first, because defer executes the defer statements in LIFO
	defer fmt.Println("Hello, World 1!")
	defer fmt.Println("Hello, World 2!")

	fmt.Println("Hello!")
	a := make(map[int]int)
	a[1] = 1
	fmt.Println(a)
}
