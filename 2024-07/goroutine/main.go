package main

import (
	"fmt"
	"sync"
)

func main() {
	exa := make(chan int)
	multi := make(chan string)

	go singleGoroutine(exa)
	go multiGoroutine(multi)

for_chan:
	for {
		select {
		case e, ok := <-exa:
			if !ok {
				break for_chan
			}
			fmt.Println(e)
		}
	}

for_multi_chan:
	for {
		select {
		case m, ok := <-multi:
			if !ok {
				break for_multi_chan
			}
			fmt.Println(m)
		}
	}

	fmt.Println("Done!")
}

func singleGoroutine(exa chan<- int) {
	go func() {
		for i := 1; i <= 10; i++ {
			exa <- i
		}
		defer close(exa)
	}()
}

func multiGoroutine(multi chan<- string) {
	go func() {
		var wg sync.WaitGroup

		for _, s := range []string{"One", "Two", "Three", "Four", "Five"} {
			wg.Add(1)
			go func() {
				defer wg.Done()
				multi <- s
			}()
		}
		wg.Wait()
		defer close(multi)
	}()
}
