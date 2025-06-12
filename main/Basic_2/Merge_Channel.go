package main

import (
	"fmt"
	"sync"
)

func streamNumbers(numbers ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, n := range numbers {
			c <- n
		}

		close(c)
	}()

	return c
}

func sumAllStreams(streams ...<-chan int) <-chan int {
	sumChan := make(chan int)
	counter := 0
	wc := new(sync.WaitGroup)

	wc.Add(len(streams))

	for i := range streams {
		go func(s <-chan int) {
			for n := range s {
				counter += n // Need to use a mutex
			}
			wc.Done()
		}(streams[i])
	}

	go func() {
		wc.Wait()
		sumChan <- counter
	}()

	return sumChan
}

func main() {
	for i := 1; i <= 100; i++ {
		s := sumAllStreams(
			streamNumbers(2, 4, 6, 8),    // Sum = 20
			streamNumbers(1, 3, 5, 7, 9), // Sum = 25
			streamNumbers(5, 10, 15, 20), // Sum = 50
		)
		x := <-s
		if x != 95 {
			fmt.Println("Error: Expected 95, got", x)
		}
	}
}
