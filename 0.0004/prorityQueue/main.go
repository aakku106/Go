package main

import (
	"fmt"
	"os"
)

var (
	queue [5][]int
	front [5]int = [5]int{-1, -1, -1, -1, -1}
)

const (
	quit = iota
	enqueue
	dequeue
	peek
)

func main() {
	fmt.Println("Trying to access prority queue.....")
	fmt.Println("Prority queue accesed")
	for {
		fmt.Println(queue)
		fmt.Println(front)
		fmt.Println("Choose Between")
		fmt.Println("1.To Enqueue")
		fmt.Println("2. To Deque")
		fmt.Println("3. Peek")
		fmt.Println("-1 To back to Queue")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case -1:
			return
		case quit:
			os.Exit(0)
		case enqueue:
			Enqueue()
		case dequeue:
			Dequeue()
		case peek:
			peekQueue()
		default:
			fmt.Println("Chooe betn 1,2,3 or 0 to exit")
		}
	}
}

func Enqueue() (int, int8, bool) {
a:

	fmt.Println("Enter the value to Enqueue in Prority Queue")
	fmt.Println("In format: value,prority")
	fmt.Println("Note: Prority shall be from 0 to 4, lower the value, higher the prority")

	var (
		value   int
		prority int8
	)

	fmt.Scanf("%d,%d", &value, &prority)

	if prority >= 5 || prority < 0 {
		fmt.Println("Prority Must be from 0 to 4")
		goto a
	}

	queue[prority] = append(queue[prority], value)

	front[prority] = front[prority] + 1
	return value, prority, false
}

func Dequeue() (int, bool) {
	if isEmpty() {
		fmt.Println("Queue is Empty !!!")
		return 0, true
	}

	for i := range 5 {
		i++
	}

	return 1, false
}

func peekQueue() {
	if isEmpty() {
		fmt.Println("Queue is Empty")
		return
	}
}
func isEmpty() bool {
	for i := range queue {

		if len(queue[i])-front[i] > 0 {
			return false
		}
	}
	return true
}
