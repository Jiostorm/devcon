package main

import (
	"fmt"
	"sync"
)

func main() {
	benchan := make(chan int)

	go query(benchan, 1)

for_chan:
	for {
		select {
		case _, ok := <-benchan:
			if !ok {
				break for_chan
			}
		}
	}

	fmt.Println("Done!")
}

func query(benchan chan<- int, frags int) {
	omega := 1_000_000 * frags

	defer close(benchan)

	var wgIn sync.WaitGroup

	for i := 1; i <= omega; i++ {
		go func() {
			wgIn.Add(1)
			defer wgIn.Done()

			benchan <- i
		}()
	}
	wgIn.Wait()
}
