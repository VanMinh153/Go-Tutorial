package main

import (
	"fmt"
	"golang.org/x/tour/tree" // Import gói tree
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}

		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println("--- Testing Walk function ---")
	ch := make(chan int)

	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	fmt.Print("Values from tree.New(1): ")
	for val := range ch {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\nExpected: 1 2 3 4 5 6 7 8 9 10")


	fmt.Println("\n--- Testing Same function ---")
	fmt.Printf("Same(tree.New(1), tree.New(1)): %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Same(tree.New(1), tree.New(2)): %v\n", Same(tree.New(1), tree.New(2)))
}