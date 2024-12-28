package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// constrained channel of size 3
	ch := make(chan int, 3)

	wg.Add(2)
	go G1(ch)
	go G2(ch)
	wg.Wait()
}

// G1 goroutine
// chan<-int implies that we can only piple values to this channel
func G1(ch chan<- int) {
	// wg is globally accessible
	defer wg.Done()

	for _, v := range []int{1, 2, 3, 4} {
		// push to the channel
		ch <- v
	}
	// in built method to close the channel
	close(ch)
}

// G2 goroutine
// chan<-int implies that we can only listen to the values from this channel
func G2(ch <-chan int) {
	// wg is globally accessible
	defer wg.Done()

	for v := range ch {
		fmt.Println(v)
	}
}
