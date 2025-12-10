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
}

func EnqueIntoLinerQueue() int {
	return 1
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
