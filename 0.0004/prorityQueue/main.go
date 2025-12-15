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
		//	fmt.Println(queue)
		//	fmt.Println(front)
		fmt.Println("Choose Between")
		fmt.Println("1.To Enqueue")
		fmt.Println("2. To Deque")
		fmt.Println("3. Peek")
		var choose int8
		fmt.Scan(&choose)
		switch choose {
		case quit:
			os.Exit(0)
		case enqueue:
			value, prority := Enqueue()
			fmt.Printf("---->\t%d\t<----enqueued with prority:\t%d\n", value, prority)
		case dequeue:
			value, wasEmpty := Dequeue()
			if !wasEmpty {
				fmt.Printf("---->\t%d\t----dequeued\n", value)
			} else {
				fmt.Println("Empty queue")
			}
		case peek:
			peekQueue()
		default:
			fmt.Println("Chooe betn 1,2,3 or 0 to exit")
		}
	}
}

func Enqueue() (int, int8) {
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

	if front[prority] == -1 {
		front[prority] = 0
	}
	// front[prority] = front[prority] + 1
	return value, prority
}

func Dequeue() (int, bool) {
	if isEmpty() {
		fmt.Println("Queue is Empty !!!")
		return 0, true
	}
	// fmt.Println("root")
	for i := range 5 {
		if front[i] > -1 {
			//	fmt.Println("i>-1")
			if front[i] == 0 {
				//	fmt.Println("f==0")
				//	fmt.Printf("index \t%d\n", i)
				front[i] = -1
				return queue[i][0], false
			} else {
				front[i] = front[i] + 1
				//	fmt.Println("f!=0")
				return queue[i][front[i-1]], false
			}
		}
	}
	return 1, true
}

func peekQueue() {
	if isEmpty() {
		fmt.Println("Queue is Empty")
		return
	}
	for i := range 5 {
		if front[i] > -1 {
			fmt.Printf("---->\t%d\t<--- will be cumming next with prority:\t%d\n", queue[i][front[i]], front[i])
			return
		}
	}
}

func isEmpty() bool {
	check := 0
	for i := range queue {
		check += len(queue[i])
	}
	if check < 0 {
		return true
	}
	return false
}
