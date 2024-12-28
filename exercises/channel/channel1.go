package main

import "fmt"

func main() {
	go func(a, b int) int {
		c := a + b
		fmt.Printf("c %v", c)
		return c
	}(1, 2)
	// Note main func will not wait till the inner func is finished
}
