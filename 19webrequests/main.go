package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.boredapi.com/api/activity"

func main() {
	// fmt.Println("LCO web request")
	response, err :=http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response is %T\n",response)
	fmt.Println(response)
	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}