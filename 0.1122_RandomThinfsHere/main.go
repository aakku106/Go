package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter How many Square checks do we have to do ?")
	fmt.Scan(&a)
	fmt.Println("You entered: ", a, " i.e. we have to check: ", a, " Numbers of time if it's square or not.")

	fmt.Println("Lect Gather numbers")
	for range a {
		var first int
		var second int
		var third int
		var forth int
		fmt.Println("Enter 1st element: ")
		fmt.Scan(&first)
		fmt.Println("Enter 2nd element: ")
		fmt.Scan(&second)
		fmt.Println("Enter 3rd element: ")
		fmt.Scan(&third)
		fmt.Println("Enter 4th element: ")
		fmt.Scan(&forth)
		if (first * 4) != (first + second + third + forth) {
			fmt.Println("Its not an square")
		} else {
			fmt.Println("It's Squeare")
		}

	}

}
