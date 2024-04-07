package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello cutie"
	file, err := os.Create("./myfile.txt")
	if err!=nil{
		panic(err)
	}
	length, err := io.WriteString(file, content)
	fmt.Println(length)
	defer file.Close()
	readFile(file)
}

func readFile(file *os.File){
	databyte, err := os.ReadFile(file.Name())
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(databyte))
}