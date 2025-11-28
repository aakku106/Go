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
}

func printSlic(s []string) {
	fmt.Printf("\ngiven Array:\t%v\tLength:\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}
