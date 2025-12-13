package main

import (
	"fmt"
)

var linearQueueStoragePlace []int
var circularQueueStoragePlace [5]int
var prorityQueueStoragePlace []int

const (
	quit = iota
	linear
	circular
	prority
)

func main() {
	for {

		fmt.Println("Welci to queue, choose between")
		fmt.Println("1. Linear Queue")
		fmt.Println("2. Circular queue")
		fmt.Println("3. Prority queue")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case quit:
			return
		case linear:
			linearQueue()
		case circular:
			circularQueue()
		default:
			fmt.Println("Choose Between 1,2,3 or choose 0 to exit")
		}
	}
}
