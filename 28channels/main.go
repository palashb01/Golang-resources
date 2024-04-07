// What is channels in go ?

// Channels are a way to communicate between goroutines.
// You can send and receive values from channels using the channel operator (<-).
// The channel is created with make() function and has a type,
// the values sent to the channel must be of the same type as the channel.
// The channel is like a pipe, you can write data to it and read data from it.
// The channel is buffered, it can hold a limited number of values,
// if the channel is full, writing to it will block until someone reads from it.

package main

import (
	"fmt"
	"sync"
)

func main(){
	// for the channel to work, in the code their should be a listener as well as sender at the same time
	myCh := make(chan int,2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// if we mark <-chan int, then it is just receive only
	// and if you are receive only then you can't have close inside the func
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChanelOpen := <-myCh
		if isChanelOpen{
			fmt.Println(val)
		}
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	// if we mark chan <-int, this means it is gonna be send only
	go func(ch chan <-int, wg *sync.WaitGroup){
		// below is how the value is added into the channel
		// myCh <- 0
		// myCh <- 6
		// close basically closes the channel
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()

}