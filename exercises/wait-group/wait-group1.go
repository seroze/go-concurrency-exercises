package main

import (
	"fmt"
	"sync"
)

func incr(wg *sync.WaitGroup) {
	var i int
	// increments the wg by 1
	wg.Add(1)
	go func() {
		// decrement the wg when it's done
		defer wg.Done()
		i++
		fmt.Println("incr go routine is finished")
	}()

}

func main() {
	var wg sync.WaitGroup
	incr(&wg)
	// Wait is like join() blocks until the wg is done
	wg.Wait()
	fmt.Println("Finished main")
}
