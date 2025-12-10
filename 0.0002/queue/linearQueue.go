package main

import (
	"fmt"
	"os"
)

const (
	enque = iota + 1
	deque
	peek
)

func linearQueue() {
	fmt.Println("Accessing Linear queue....")
	fmt.Println("Linear queue accesed")
	for {
		fmt.Printf("choose between:\n1.\tEnque\n2.\tDeque\n3.\tPeek\n")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case quit:
			os.Exit(0)
		case enque:
			fmt.Printf("--->\t%d\tInserted in Linear Queue", EnqueIntoLinerQueue())
		case deque:
			dequedValueFromLinearQueue, wasLinearQueueEmpty := DequFromLinerQueue()
			if wasLinearQueueEmpty {
				fmt.Println("Linear Queue was empty...")
			} else {
				fmt.Printf("---->\t%d\t<---- Dequed", dequedValueFromLinearQueue)
			}
		case peek:
			PeekIntoLinearQueue()
		default:
			fmt.Println("Choose betn 1,2,3 or choose 0 to exit program")
		}
	}
}

func EnqueIntoLinerQueue() int {
	fmt.Println("Enter value to enqueue")
	var valueToEnqueueInLinearQueuesTop int
	fmt.Scan(&valueToEnqueueInLinearQueuesTop)
	linearQueueStoragePlace = append(linearQueueStoragePlace, valueToEnqueueInLinearQueuesTop)
	return valueToEnqueueInLinearQueuesTop
}
func DequFromLinerQueue() (int, bool) {
	if isLinearQueueEmpty() {
		return 0, false
	}
	dequedValue := linearQueueStoragePlace[0]
	linearQueueStoragePlace = linearQueueStoragePlace[1:]
	return dequedValue, true
}
func isLinearQueueEmpty() bool {
	if len(linearQueueStoragePlace) <= 0 {
		return true
	}
	return false
}
func PeekIntoLinearQueue() {
	if isLinearQueueEmpty() {
		fmt.Println("The Linear Queue is Empty !!!")
	} else {
		fmt.Printf("----->\t%d\t<---- is next to come out", linearQueueStoragePlace[0])
	}
}
