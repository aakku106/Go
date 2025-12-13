package main

import (
	"fmt"
	"os"
)

var frontOFcirularQueue, rearOfCircularQueue int = -1, -1

func circularQueue() {
	fmt.Println("Accessing Circular queue....")
	fmt.Println("Circular queue accesed")
	for {
		//	fmt.Printf("\n\n\tqueue:\t%v\tfront:\t%d\trear\t%d\tlen:%d\tcap:%d\n", circularQueueStoragePlace, frontOFcirularQueue, rearOfCircularQueue, len(circularQueueStoragePlace), cap(circularQueueStoragePlace))
		fmt.Printf("choose between:\n1.\tEnque\n2.\tDeque\n3.\tPeek\n-1. Back\n")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case back:
			return
		case quit:
			os.Exit(0)
		case enque:
			enqueuedValueInCircularQueue, wasCircularQueueFull := EnqueueInCircularQueue()
			if !wasCircularQueueFull {
				fmt.Printf("--->\t%d\tInserted in Circular Queue\n", enqueuedValueInCircularQueue)
			} else {
				fmt.Println("The Circular Queue is Full")
			}
		case deque:
			dequedValueFromCircularQueue, wasCircularQueueEmpty := DequeFromCircularQueue()
			if !wasCircularQueueEmpty {
				fmt.Printf("---->\t%d\t<---- Dequed", dequedValueFromCircularQueue)
			} else {
				fmt.Println("Circular Queue was empty...")
			}
		case peek:
			PeekIntoCircularQueue()
		default:
			fmt.Println("Choose betn 1,2,3 or choose 0 to exit program")
		}
	}
}

func EnqueueInCircularQueue() (int, bool) {
	if !isCircularQueueFull() {
		var value int
		fmt.Println("Give data to enqueue in circular Queue")
		fmt.Scan(&value)
		// rearOfCircularQueue++
		rearOfCircularQueue = (rearOfCircularQueue + 1) % len(circularQueueStoragePlace)
		circularQueueStoragePlace[rearOfCircularQueue] = value
		if frontOFcirularQueue < 0 {
			frontOFcirularQueue, rearOfCircularQueue = 0, 0
		}
		return value, false
	}
	fmt.Println("The Circular Queue is Full")
	return 0, true
}
func DequeFromCircularQueue() (int, bool) {
	if !isCircularQueueEmpty() {
		dequeuedValue := circularQueueStoragePlace[frontOFcirularQueue]
		// frontOFcirularQueue++
		if frontOFcirularQueue == rearOfCircularQueue {
			frontOFcirularQueue, rearOfCircularQueue = -1, -1
		} else {
			frontOFcirularQueue = (frontOFcirularQueue + 1) % len(circularQueueStoragePlace)
		}
		return dequeuedValue, false
	}
	return 0, true
}

func PeekIntoCircularQueue() {
	if isCircularQueueEmpty() {
		fmt.Println("The Circular Queue is empty")
	} else {
		fmt.Printf("---->%d<----will come out next...\n", circularQueueStoragePlace[frontOFcirularQueue])
	}
}
func isCircularQueueEmpty() bool {
	if frontOFcirularQueue < 0 {
		return true
	}
	return false
}
func isCircularQueueFull() bool {
	if ((rearOfCircularQueue + 1) % len(circularQueueStoragePlace)) == frontOFcirularQueue {
		return true
	}
	return false
}
