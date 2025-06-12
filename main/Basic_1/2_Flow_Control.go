package main

import (
	"fmt"
	// "math"
	// "os"
	// "runtime"
	"time"
)

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

func main() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// sum := 1
	// for ; sum < 1000; {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// for {}

	// fmt.Println(sqrt(2), sqrt(-4))

	// if i := 2520; i%7 == 0 {
	// 	fmt.Println("i is divisible by 7")
	// } else {
	// 	fmt.Println("i is not divisible by 7")
	// }

	// fmt.Print("Go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("macOS.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }

	// fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	// fmt.Println("Today is", today)
	// fmt.Println("Saturday is", time.Saturday)

	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("VSL")
	fmt.Print("Welcome to ")

}
