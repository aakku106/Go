package main

import "fmt"

type cat interface{}

// this is how any keywoard is made, actually any is not even keywoard
// its interface with no methods
//
//	In Befores example given for calculation of area of different sizes, i used (name *shape name), but we dont always have to do that in interface
type shape interface {
	set(...float64)
	area() float64
}
type rectangle struct{ length, breadth float64 }

func printArea(s shape) { fmt.Println("Area: ", s.area()) }

// func printArea(s shape)           { fmt.Println("Area: ", s.area()) }
// func (r rectangle) area() float64 { return r.breadth * r.length }
func (r *rectangle) set(nums ...float64) {
	switch {
	case len(nums) == 2:
		r.length = nums[0]
		r.breadth = nums[1]
	default:
		fmt.Println("Not enough values")
	}
}
func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func main() {
	newRect := rectangle{}
	go newRect.set(123.321, 321.123)
	fmt.Println(newRect.area()) // Here i printed area in main function, i can insted use iinterface
	printArea(&newRect)
}
