package main

import (
	"fmt"
	// "math/cmplx"
)

// Functions
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (y, x int) {
	x = sum % 10
	y = sum - x
	return
}

// Variables

// Main
func main() {
	// fmt.Println(add(42, 13))

	// a, b := swap("Go", "Language")
	// fmt.Println(a, b)
	// fmt.Println(swap(swap("Go", "Language")))

	// fmt.Println(split(19))

	// var i, j int = 22, 66
	// x, y, z := true, false, "Oke!"

	// fmt.Println(i, j, x, y, z)

	// var (
	// 	v1 bool       = false
	// 	v2 uint64     = 1<<64 - 1
	// 	v3 complex128 = cmplx.Sqrt(-5 + 12i)
	// )

	// fmt.Printf("Type: %-18T Value: %v\n", v1, v1)
	// fmt.Printf("Type: %-18T Value: %v\n", v2, v2)
	// fmt.Printf("Type: %-18T Value: %v\n", v3, v3)

	// i := 299
	// j := float64(i)
	// k := uint(j)
	// fmt.Println(i, j, k)

	// const Pi = 3.14
	// fmt.Println("Happy", Pi, "Day")

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(float64(Big))
	fmt.Println(Small)
}
