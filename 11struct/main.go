package main

import "fmt"

func main(){
	fmt.Println("structs in golang")
	hitesh := User{"Hitesh", "hitesh@go.dev",true,32}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	// no inheritance in golang
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}