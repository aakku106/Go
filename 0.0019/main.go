// This file is made to see very basics of defer function in go and has no test files
// Here we will understand defer from very basics
package main

import "fmt"

func main() {
	defer fmt.Println("defer print")
	fmt.Println("Normal Print")
	/*
		Wired right ever tho
			defer fmt.Println("defer print")
		was before it printed after
			fmt.Println("Normal Print")
			/*
	*/
}

// Lest explore further in ./next/defer.go
