package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main entered")
	// unconstrained channel
	ch := make(chan string)

	go func() {
		time.Sleep(5 * time.Millisecond)
		ch <- "message"
	}()

	select {
	case m := <-ch:
		fmt.Println("received message: ", m)
	default:
		// when is this useful ?
		// this is like a default flow
		fmt.Println("No message received")
	}

}
