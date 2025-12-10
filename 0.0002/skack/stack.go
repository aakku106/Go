package main

import (
	"fmt"
)

const (
	exit = iota
	insert
	remove
	peek
)

func stack() int {
	for {

		fmt.Println("Accessing stack")
		fmt.Println("Stack accessed, choose between")
		fmt.Printf("\n1. Insert\n2. Remove\n3. Peek\n")
		choose := -1
		fmt.Scan(&choose)

		switch choose {
		case exit:
			return 0
		case insert:
			fmt.Printf("--->\t%d\tInserted in stack", InsertInStack())
		case remove:
			tempPopedValue, isNotEmpty := RemoveFromStack()
			if isNotEmpty {
				fmt.Printf("\t%d\tpoped from stack", tempPopedValue)
			} else {
				fmt.Println("Stack is empty")
			}

		case peek:
			PeekInStack()
		default:
			fmt.Println("---Choose between 0,1,2,3")
		}
	}
}
func InsertInStack() int {
	var value int
	fmt.Println("Enter the number ot put on stack")
	fmt.Scan(&value)
	arr = append(arr, value)
	// notice you cant do ++ pointer in go
	// Go people say it's to remove ambiguity,but i say i am more confused now xd
	return value
}
func RemoveFromStack() (int, bool) {
	if !isStackEmpty() {
		// fmt.Printf("\t%d\t%v\t", len(arr), arr)
		lastIndexValue := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		//	fmt.Printf("\t%d\t%v\t", len(arr), arr)
		return lastIndexValue, true
	}
	return 0, false
}

func PeekInStack() {
	if !isStackEmpty() {
		fmt.Printf("arr: %d\n", arr[len(arr)-1])
	} else {
		println("Stack is Empty")
	}
}
func isStackEmpty() bool {
	// fmt.Printf("--------\t%d\t-------------", len(arr))
	if len(arr) <= 0 {
		return true
	}
	return false
}
