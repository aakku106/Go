package main

import "fmt"

func main() {
	a := make([]int, 0, 4)
	fmt.Println(a)
	a = append(a, 12)
	fmt.Println(a)
	b := make([]int, 5)
	fmt.Println(b)
	b = append(b, 123)
	fmt.Println(b)
}
