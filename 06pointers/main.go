package main

import "fmt"

func main(){
	fmt.Println("Hello, World!");
	mynum := 23
	var ptr = &mynum
	fmt.Println(ptr)
}