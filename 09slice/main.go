package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to video on slices")
	var fruitList = []int{1,2,3}
	fmt.Println("Fruit List is: ", fruitList)
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, 4, 5, 6)
	fruitList = append(fruitList[1:4])
	fmt.Println("Fruit List after append is: ", fruitList)
	fruitList = append(fruitList[:2], fruitList[3:]...)
	fmt.Println("Fruit List after deleting is: ", fruitList)
	highScores := make([]int, 2)
	highScores[0] = 234
	highScores[1] = 945
	// highScores[2] = 999 // This will give an error as we are trying to access a non-existing index
	// but with the append this works because when we use append we overwrite everything in a new memory add.
	highScores = append(highScores, 555, 666, 777)
	// fmt.Println("High Scores: ", highScores)
	sort.Ints(highScores)
	fmt.Println("Sorted High Scores: ", highScores)
	// the below code is for removing particular index from the slice
	var index int = 2;
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
} 