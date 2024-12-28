package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]string)

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()

		for len(sharedRsc) == 0 {
			// while sharedRsc is empty
			// unlock it sleep for 1 ms by then the main thread might have
			// updated the sharedRsc and then locked it
			mu.Unlock()
			time.Sleep(1 * time.Millisecond)
			mu.Lock()
		}

		fmt.Println(sharedRsc["rsc1"])
		mu.Unlock()
	}()

	mu.Lock()
	sharedRsc["rsc1"] = "foo"
	mu.Unlock()

	wg.Wait()
}
