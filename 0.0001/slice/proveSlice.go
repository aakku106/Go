package main

import "fmt"

func proveSlice() {
	// What i gonna show here ?
	// - Slice do not own the array, it is just pointer to underlying array in memmory

	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------Prove slice-----------------")
	arr := []string{} //capacity 0 simple
	printSlic(arr)
	arr = append(arr, "Cats") //capacity 1 simple
	printSlic(arr)
	arr = append(arr, "Eats") //capacity 2 simplee
	printSlic(arr)
	arr = append(arr, "Rats") //Capacity 4 wtf
	printSlic(arr)
	// lest do one more time than you will understand the pattern
	arr = append(arr, "and") // see, now capacity is still 4 and length is also 4.
	printSlic(arr)
	arr = append(arr, "Becomes") // Here capacity becomes 8 and lenght is 5
	printSlic(arr)
	arr = append(arr, "Fat") // Capacity is still 8 and length gets 6
	printSlic(arr)
	// so, Go simplly doubles the size of array or slice when its capacity gets filled,
	// Again slice is a array with fixed capacity, go just doubles the size of array
}

func printSlic(s []string) {
	fmt.Printf("\ngiven Array:\t%v\tLength:\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}
