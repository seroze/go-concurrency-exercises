package main

import (
	"fmt"
	"time"
)

func main() {

	var data int
	go func() {
		data++
	}()
	time.Sleep(10 * time.Millisecond)
	if data == 0 {
		fmt.Println("value of data is 0")
	} else if data == 1 {
		fmt.Println("value of data is 1")
	}
}
