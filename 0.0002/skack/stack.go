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
		temp := RemoveFromStack()
		if temp != -106 {
			fmt.Printf("\t%d\tpoped from stack", temp)
		} else {
			fmt.Println("Stack is empty")
		}

	case peek:
		Peek()
	default:
		fmt.Println("---Choose between 0,1,2,3")
	}
	stack()
	return 0
}
func InsertInStack() int {
	var value int
	fmt.Println("Enter the number ot put on stack")
	fmt.Scan(&value)
	arr = append(arr, value)
	pointer++
	// notice you cant do ++ pointer in go
	// Go people say it's to remove ambiguity,but i say i am more confused now xd
	return value
}
func RemoveFromStack() int {
	if isStackEmpty() {
		// fmt.Printf("\t%d\t%v\t", len(arr), arr)
		lastIndexValue := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		//	fmt.Printf("\t%d\t%v\t", len(arr), arr)
		return lastIndexValue
	}
	return -106
}

func Peek() {
	if isStackEmpty() {
		fmt.Printf("arr: %d\n", arr[len(arr)-1])
	} else {
		println("Stack is Empty")
	}
}
func isStackEmpty() bool {
	// fmt.Printf("--------\t%d\t-------------", len(arr))
	if len(arr) <= 0 {
		return false
	}
	return true
}
