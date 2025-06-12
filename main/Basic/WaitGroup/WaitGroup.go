package main

import (
	"fmt"
	"sync"
)

// func hello(wgrp *sync.WaitGroup) {
// 	fmt.Println("Hello")
// 	wgrp.Done()
// }

func main() {
	//     var wg sync.WaitGroup
	//     wg.Add(2)
	//     go hello(&wg)
	//     go hello(&wg)
	//     wg.Wait()
	//     fmt.Println("main function")

	var wg sync.WaitGroup
	players := []string{"James Harden", "Kyrie Irving", "Kevin Durant"}
	wg.Add(len(players))
	for _, player := range players {
		go func(baller string) {
			fmt.Printf("printing player %s\n", baller)
			wg.Done()
		}(player)
	}
	wg.Wait()

}
