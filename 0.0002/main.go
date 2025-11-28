package main

import "fmt"

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------Error Handling'Level: Easy' -----------------")
	// You better look at 0.0001/slice/proveSlice to understand whats going on here
	arr := []string{}
	printSlic(arr)
	arr = append(arr, "Cats")
	printSlic(arr)
	arr = append(arr, "Eats")
	printSlic(arr)
	arr = append(arr, "Rats")
	printSlic(arr)
	arr = append(arr, "and")
	printSlic(arr)
	arr = append(arr, "Becomes")
	printSlic(arr)
	arr = append(arr, "Fat")
	printSlic(arr)
	newArr := arr[:]
	printSlic(newArr)
	newArr[5] = "Super Fat"
	printSlic(newArr)
	printSlic(arr)
	newArr[7] = "last index" // we will get runtime error here
	if newArr[7] == nil {
		println("cant access 7 space its out of length of slice ")
	}
}

func printSlic(s []string) {
	fmt.Printf("\ngiven Array:\t%v\tLength:\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}

