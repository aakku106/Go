package main

import (
	"fmt"
	"os"
)

var (
	frontOfProrityQueue int = -1
)

func prorityQueue() {
	fmt.Println("Accessing Protity Queue.....")
	fmt.Println("Protuty Queue Accesed")
	fmt.Println("Choose Between")
	fmt.Println("1.To Enqueue")
	fmt.Println("2. To Deque")
	fmt.Println("3. Peek")
	var choose int8
	fmt.Scan(&choose)
	switch choose {
	case enque:
		enqueuedValueInProrityQueue, prorityOfEnqueuedValueInProrityQueue := EnqueueInProrityQueue()
		fmt.Printf("---->\t%d\t<----Enqueued Into Prority Queue with Prority of:\t%d", enqueuedValueInProrityQueue, prorityOfEnqueuedValueInProrityQueue)
	case deque:
	case peek:
		peekIntoProrityQueue()
	default:
		fmt.Println("Choose between 1,2,3 or 0 to exit program")
	}

}
func EnqueueInProrityQueue() (int, int) {
	return 1, 1
}
func peekIntoProrityQueue() {
	if isProrityQueueEmpty() {
		fmt.Println("The Prority Queue is Empty")
		return
	}
	fmt.Printf("---->\t%d\t<---- Will come out next", prorityQueueStoragePlace[frontOfProrityQueue])
}
func isProrityQueueEmpty() bool {
	if frontOFcirularQueue == -1 {
		return true
	}
	return false
}

/*
Lest see how we make  prority queue:
1. 1st it shall follow fifo rele: first in 1st out,which means i have to 1st make a normal linear queue(not circular queue cause i have a idea to use slice like a ummm some shord of garbage collector thing like, when lest say slice is larger than 50, ie. the front is larger than 50 than only do [front:], and front=0; i think it should be ~ O(1)may be lest cee)
2. i need a prority, so lest decleare a var called prority, where we can give values from 0 to 4, lower the value higher the prority, I think this is good idea for now
*/
