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
	if 7 >= len(newArr) {
		fmt.Println("slice newArrar only has length of", len(newArr), "So you can't access out of ", len(newArr)-1)
	} else {
		newArr[7] = "last Index"
	} // wait were you aspectign == nil thing just now lol xd
	// Next go to pacakages/main.go, there is more intresting stuff
}

func printSlic(s []string) {
	fmt.Printf("\ngiven Array:\t%v\tLength:\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}
