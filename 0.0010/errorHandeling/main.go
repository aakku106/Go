package main

import "fmt"

func main() {
	if ok, err := isEven(10); ok {
		fmt.Println("It's even")
	} else {
		fmt.Println(err)
	}
}

func isEven(n int) (bool, error) {
	if n&1 == 1 {
		return false, fmt.Errorf("It's Odd")
	}
	return true, nil
}
