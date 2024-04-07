package main

import (
	"encoding/json"
	"fmt"
)

type courses struct{
	// because of this json: coursename when the json will be sent it will be renamed to coursename
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	// because of the '-' this will never be reflected in the json
	Password string `json:"-"`
	// because of the omitempty, if the tags is empty for a particular data, it will not be shown
	Tags []string	`json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCources := []courses{
		{"ReactJS Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"Angular Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"VueJS Bootcamp",299,"LearnCodeOnline.in","abc123",nil},
	}
	finalJSON,_ := json.MarshalIndent(lcoCources,"","\t")
	fmt.Println(string(finalJSON))
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"website": "learncodeonline.in",
			"tags": ["web-dev","js"]
		}
	`)
	// var lcoCources courses

	// checkValid := json.Valid(jsonDataFromWeb)
	// if checkValid{
	// 	fmt.Println("JSON was valid")
	// 	json.Unmarshal(jsonDataFromWeb,&lcoCources)
	// 	fmt.Printf("%#v\n",lcoCources)
	// }else{
	// 	fmt.Println("JSON was not valid")
	// }

	// When the data is coming from the web, we are not sure of the value type, we are sure of the string, but no the value so because of which we use interface.
	var myOnlineData map[string]interface{}

	err := json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%#v\n",myOnlineData)
	for k,v := range myOnlineData{
		fmt.Printf("key is %v and value is %v and type is %T\n",k,v,v)
	}

}