package main

import (
	"fmt"
	"runtime"
)

var (
	c1, c2, c3 int = 0, 0, 0
)

func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("1 has sent by %s\n", name)
		switch name {
		case "S1":
			c1++
			break
		case "S2":
			c2++
			break
		case "S3":
			c3++
			break
		}
		runtime.Gosched()
	}
}

func main() {
	myChan := make(chan int)

	go sender(myChan, "S1")
	go sender(myChan, "S2")
	go sender(myChan, "S3")

	start := 0

	for {
		start += <-myChan
		fmt.Println(start)

		if start >= 300 {
			break
		}
		// runtime.Gosched()
	}

	fmt.Printf("S1: %d, S2: %d, S3: %d\n", c1, c2, c3)
}
