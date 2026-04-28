package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println(f(5000))
	})

	wg.Go(func() {
		fmt.Println(ff(20))
	})
	wg.Wait()
}

func f(n int) ([]int, error) {
	if n <= 1 {
		return nil, fmt.Errorf("Cant be smaller than 2")
	}

	arr := make([]int, n)
	arr[0], arr[1] = 0, 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr, nil
}
func ff(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return ff(n-1) + ff(n-2)
}
