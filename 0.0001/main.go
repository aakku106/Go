package main

import "fmt"

func main() {
	fmt.Println("i like Cat")
	a := 1
	b := 2
	c := a + b

	fmt.Println(a, b, c)
	cat := "a domestic animal"
	bol := true
	rat := "eat it"
	fmt.Println(cat, rat, bol)

	// array

	arr := [5]int{1, 2, 3, 4, 5}
	anotherArray := [2]string{"cat", "rat"}
	fmt.Println(arr, anotherArray, anotherArray[0], anotherArray[1], arr[4])

	// Slice
	// thay are like brothers ot array, but more opean minded xd..

	sli := []int{1, 2, 3}
	fmt.Println(sli)
	sli = append(sli, 4)
	fmt.Println(sli)
	sli = append(sli, 4)
	fmt.Println(sli)
	NewSlice := sli[1:3]
	fmt.Println(NewSlice)
	// this is somewhat familiar to Python ...

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)
	ss := s[:0]
	printSlice(ss)
	sss := s[1:]
	printSlice(sss)
	// of curse we can reuse a variable
	sss = sss[2:]
	printSlice(sss)
	sss = sss[:2]
	printSlice(sss)
	sss = sss[1:2]
	printSlice(sss)
	// notice the capacity

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
