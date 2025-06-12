package main

import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	fmt.Println("sum is", sum)
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	// pos, neg := adder(), adder()
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }

	f := fibonacci()
	for i := 1; i <= 40; i++ {
		fmt.Printf("%2d: %-d\n", i, f())
	}
}
