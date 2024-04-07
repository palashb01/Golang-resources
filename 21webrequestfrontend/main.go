package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const url string = "http://localhost:1111/get"
	response, err :=http.Get(url)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	dataBytes, err := io.ReadAll(response.Body)
	var responseString strings.Builder
	byteCount,_ := responseString.Write(dataBytes)
	fmt.Println(byteCount)
	fmt.Println(responseString)
	fmt.Println(responseString.String())
	// fmt.Println(string(dataBytes))
}

func PerformPostJsonRequest(){
	const url string = "http://localhost:1111/post"
	jsonData := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "youtube"
		}
	`)
	response, err := http.Post(url, "application/json", jsonData)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	dataBytes, err := io.ReadAll(response.Body)
	fmt.Println(string(dataBytes))
}


func PerformPostFormRequest(){
	const myurl string = "http://localhost:1111/postform"
	data:= url.Values{}
	data.Add("firstName","Palash")
	data.Add("","hitesh")
	response, err := http.PostForm(myurl, data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
}