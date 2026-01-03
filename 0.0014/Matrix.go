package main

import (
	"fmt"
	"sync"
)

func main() {

	easyWay()
	// This looks no fun what if you could do
	//	fmt.Println(a - b)
	// But this isen't matlab yet,so
	// So lest explore given libraryes
	// It seems by defult Go peope dont ship these kind of code/libraries at all
	// So what we could do is
	WaitWay()
	// But it will cost more than easyWay cause subtraction is just 1 instruction to CPU not parellel code

}
func easyWay() {
	a := []uint{2, 2, 2, 2, 2}
	b := []uint{1, 1, 1, 1, 1}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i] - b[i])
	}
}
func WaitWay() {
	a := []uint{2, 2, 2, 2, 2}
	b := []uint{1, 1, 1, 1, 1}

	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(a[i] - b[i])
		}(i)
	}

	wg.Wait()
}
