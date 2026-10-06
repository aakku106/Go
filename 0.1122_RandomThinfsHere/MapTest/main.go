package main

import "fmt"

func main() {

	a := make(map[int]int, 5)

	a[1] = 2
	a[1] = 3
	a[2] = 23

	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)

	}

}
