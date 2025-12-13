package main

import (
	"fmt"
)

var (
	linearQueueStoragePlace   []int
	circularQueueStoragePlace [5]int
	prorityQueueStoragePlace  [5][]int
)

const (
	quit = iota
	linear
	circular
	prority
)
const (
	back  = -1
	enque = iota + 1
	deque
	peek
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
		case prority:
			prorityQueue()
		default:
			fmt.Println("Choose Between 1,2,3 or choose 0 to exit")
		}
	}
}
