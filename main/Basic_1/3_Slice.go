package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice("s1", s)

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice("s2", s)

	// // Extend its length.
	// s = s[:4]
	// printSlice("s3", s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice("s4", s)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
