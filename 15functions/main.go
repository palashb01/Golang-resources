package main

import "fmt"

func main(){
	fmt.Println("Welcome to functions in golang")
	// greeter()
	// nested functions are not allowed
	result:=adder(1,2)
	results,_:=proAdders(1,2)
	fmt.Println("Result is: ",result)
	fmt.Println(results)
}

// indefinite arguments when you don't know the number of arguments
func proAdder(values ...int) int{
	total:=0
	for _,values:=range values{
		total+=values
	}
	return total
}

// to return multiple values
func proAdders(values ...int) (int,string){
	total:=0
	for _,values:=range values{
		total+=values
	}
	return total,"hi"
}


// below line is called function signature
func adder(a int, b int) int{
	return a+b
}

func greeter(){
	fmt.Println("Namaste from golang")
}