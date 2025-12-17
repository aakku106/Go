package main

import "fmt"

func main() {
	// factorial()
	fiboNachi()

}
func factorial() {
	num := 5
	for i := 1; i < 5; i++ {
		num *= i
	}
	fmt.Println(num)
}
func fiboNachi() {
	val := 1
	for i := range 10 {
		val += i
		// fmt.Print(" ", val, " ")
		fmt.Printf("\t%d + %d\n", val, i)
	}
}
