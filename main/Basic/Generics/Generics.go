package main

import (
	"fmt"
)

func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

type SignedInteger interface {
  int | int8 | int16 | int32 | int64
}

func Smallest[T SignedInteger](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r > v {
			r = v
		}
	}
	return r
}

func main() {
	arr := []int{1, 2, 3}
	resultArr := Map(arr, func(v int) int { return v * 2 })
	fmt.Println(resultArr)

	arr2 := []string{"a", "b", "c"}
	resultArr2 := Map(arr2, func(v string) string { return v + v })
	fmt.Println(resultArr2)

	ints := []int{2, 4, 6}
	smallestInt := Smallest(ints)
	fmt.Println(smallestInt)
}

// [2, 4, 6]
// ["aa", "bb", "cc"]
