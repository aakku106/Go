package main

import "fmt"

func main() {
	a := 3
	b := 3
	fmt.Println(add(a, b))
	fmt.Println(add(a, b, 3))

}

func add(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s

}
