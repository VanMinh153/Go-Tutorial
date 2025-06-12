package main

import (
	"fmt"
)

func main() {
	var arr []int = []int{5, 10, 15, 20, 25}
	for _, v := range arr {
		fmt.Println("Value:", v)
	}

	c := make(chan int, 5)
	c <- 5
	c <- 10
	c <- 15
	close(c)
	for i := range c {
		fmt.Println("Received:", i)
	}
}
