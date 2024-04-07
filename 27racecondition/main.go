package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait group is necessary to that roroutines does not exit
	// mutex is imp to prevent the race condition
	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
