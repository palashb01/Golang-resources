package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	var arr = [3]int{1,2,3}
	fmt.Println(arr);
	for i:=0; i<len(arr);i++{
		fmt.Println(arr[i])
	}
}