package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Learning sync waitgroup in golang")
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1) // add 1 to the waitgroup
		go worker(i, &wg)
	}

	wg.Wait()
}

func worker(i int, wg *sync.WaitGroup) {

	defer wg.Done() // decrement the waitgroup
	fmt.Printf("worker %d starting\n", i)
	// do some work
	fmt.Printf("worker %d done\n", i)
}
