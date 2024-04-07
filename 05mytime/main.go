package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	fmt.Println("Hello, World!");
	presentTime:=time.Now()
	fmt.Println(presentTime.Format("01-02-2006 03:04:05 Monday"))
	fmt.Println(presentTime)
	created:= time.Date(2020, time.August,23,22,22,22,22,time.UTC);
	fmt.Println(created.Format("01-02-2006 Monday"))
	fmt.Println(runtime.NumCPU())
}