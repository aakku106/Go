package main

import "fmt"

func main() {
	var a [2]byte
	var b [2]byte

	c := make(map[[2]byte]int)
	a[0] = 0
	b[0] = 1
	a[1] = 1
	b[1] = 0

	c[a] = 1
	c[b] = 2

	fmt.Println(c)

}
