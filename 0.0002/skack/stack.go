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
		fmt.Printf("%d Inserted in stack", InsertInStack())
	case remove:
		RemoveFromStack()
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
	return value
}
func RemoveFromStack() int {
	val := 2
	return val
}

func Peek() {
	fmt.Printf("arr: %v\n", arr[pointer])
}
