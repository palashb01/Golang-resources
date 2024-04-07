package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, World!")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":1111",r))
}

func greeter(){
	fmt.Println("Hey there")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my server</h1>"))
}