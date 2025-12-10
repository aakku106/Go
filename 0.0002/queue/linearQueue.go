package main

import "fmt"

const (
	enque = iota + 1
	deque
	peek
)

func linearQueue() {
	fmt.Println("Accessing Linear queue....")
	fmt.Println("Linear queue accesed")
	fmt.Printf("choose between:\n1.\tEnque\n2.\tDeque\n3.\tPeek")
	var choose int8
	fmt.Scan(&choose)
	switch choose {

	}
}

func EnqueIntoLinerQueue() int {
	fmt.Println("Enter value to enqueue")
	var valueToEnqueueInLinearQueuesTop int
	fmt.Scan(&valueToEnqueueInLinearQueuesTop)
	return valueToEnqueueInLinearQueuesTop
}
func DequFromLinerQueue() (int, bool) {
	if isLinearQueueEmpty() {
		return 0, false
	}
	return 1, true
}
func isLinearQueueEmpty() bool {
	if len(linearQueueStoragePlace) <= 0 {
		return true
	}
	return false
}
