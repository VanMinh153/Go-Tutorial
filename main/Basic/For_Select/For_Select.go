package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 1; i < 40; i++ {
			fmt.Printf("%2d: %-d\n", i, <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
