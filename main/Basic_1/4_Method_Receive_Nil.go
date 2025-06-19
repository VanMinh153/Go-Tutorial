package main

import "fmt"

type MyInterface interface {
	M()
}

type T struct {
	S string
}

// Phương thức M cho kiểu *T
func (t *T) M() {
	if t == nil {
		fmt.Println("nil receiver (t is nil)")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i MyInterface // i là nil interface (value=nil, type=nil)
	fmt.Printf("Initial i: (value=%v, type=%T)\n", i, i) // Output: Initial i: (value=<nil>, type=<nil>)

	var t *T // t là con trỏ nil của kiểu T (value=nil, type=*T)
	fmt.Printf("Initial t: (value=%v, type=%T)\n", t, t) // Output: Initial t: (value=<nil>, type=*main.T)

	i = t // Gán con trỏ nil t vào interface i
	fmt.Printf("After assigning t to i: (value=%v, type=%T)\n", i, i) // Output: After assigning t to i: (value=<nil>, type=*main.T)

	i.M() // Gọi phương thức M trên interface i
}