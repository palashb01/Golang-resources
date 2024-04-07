package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	// below go keyword is used to create a new go routine
// 	// if you just run this without sync package then hello will not be printed
// 	// reason:
// 	// The reason "hello" might not be printed is due to the fact that the main goroutine (the one running main()) may exit before the goroutine for "hello" has a chance to execute. Since goroutines run concurrently, there's no guarantee that the "hello" goroutine will start or finish before the main goroutine completes. When the main goroutine exits, it also ends the program, which can halt the execution of any remaining goroutines
// 	go greeter("hello")
// 	greeter("world")
// }

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

var signals = []string{"test"}

// usually these are pointers
var wg sync.WaitGroup
var mut sync.Mutex

func main(){
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _,web:= range websitelist{
		go getStatusCode(web)
		wg.Add(1)
	}
	// the below line should always be the last line
	wg.Wait()
	fmt.Println(signals)
}
func getStatusCode(endpoint string){
	defer wg.Done()
	res,err:=http.Get(endpoint)
	if err!=nil{
		fmt.Println("Error: ",err)
	}
	mut.Lock()
	signals = append(signals,endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
}
