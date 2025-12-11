package main

import (
	"fmt"
	"os"
)

var frontOFcirularQueue, rearOfCircularQueue int = 0, 0

func circularQueue() {
	fmt.Println("Accessing Circular queue....")
	fmt.Println("Circular queue accesed")
	for {
		fmt.Printf("choose between:\n1.\tEnque\n2.\tDeque\n3.\tPeek\n")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case quit:
			os.Exit(0)
		case enque:
			enqueuedValueInCircularQueue, wasCircularQueueFull := EnqueueInCircularQueue()
			if wasCircularQueueFull {
				fmt.Println("The Circular Queue is Full")
			}
			fmt.Printf("--->\t%d\tInserted in Circular Queue\n", enqueuedValueInCircularQueue)
		case deque:
			dequedValueFromCircularQueue, wasCircularQueueEmpty := DequFromLinerQueue()
			if wasCircularQueueEmpty {
				fmt.Println("Linear Queue was empty...")
			} else {
				fmt.Printf("---->\t%d\t<---- Dequed", dequedValueFromCircularQueue)
			}
		case peek:
			PeekIntoCircularQueue()
		default:
			fmt.Println("Choose betn 1,2,3 or choose 0 to exit program")
		}
	}
}

func EnqueueInCircularQueue() int {
	return 1
}

func PeekIntoCircularQueue() {
	if isCircularQueueEmpty() {
		fmt.Println("The Circular Queue is empty")
	} else {
		fmt.Printf("---->%d<----will come out next...\n", 1)
	}
}
func isCircularQueueEmpty() bool {
	if len(circularQueueStoragePlace) <= 0 {
		return true
	}
	return false
}
