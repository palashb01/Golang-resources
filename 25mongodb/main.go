package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("mongodb API")
	fmt.Println("server is running")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
}