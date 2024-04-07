package main

import "fmt"

func main(){
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	// keys := make([]string, 0, len(languages))
	// for k := range languages {
	// 	keys = append(keys, k)
	// }
	// loops
	for _, value := range languages{
		fmt.Println(value)
	}
}