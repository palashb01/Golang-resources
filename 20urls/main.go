package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://datausa.io/api/data?drilldowns=Nation&measures=Population"

func main() {
	// fmt.Println("Welcome to handling URLs in golang")
	result,_ := url.Parse(myurl)
	// fmt.Println(result)
	// fmt.Println(result.Host)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n",qparams)
	fmt.Println(qparams["drilldowns"])
	for _,values := range qparams{
		fmt.Println(values)
	}

	// This is to create URL
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "datausa.io",
		Path: "/api/data",
		RawQuery: "drilldowns=Nation&measures=Population",
	}
	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}