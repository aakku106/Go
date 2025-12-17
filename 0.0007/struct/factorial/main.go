package main

import "fmt"

func main() {
	num := 5
	for i := 1; i < 5; i++ {
		num *= i
	}
	fmt.Println(num)
}
