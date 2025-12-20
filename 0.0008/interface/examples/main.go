package main

import "fmt"

type cat interface{}

// this is how any keywoard is made, actually any is not even keywoard
// its interface with no methods
//  In Befores example given for calculation of area of different sizes, i used (name *shape name), but we dont always have to do that in interface

type rectangle struct{ length, breadth float64 }
type shape interface {
	set(...float64)
	area() float64
}

func printArea(s shape)           { fmt.Println("Area: ", s.area()) }
func (r rectangle) area() float64 { return r.breadth * r.length }
func set(nums ...float64) {

}

func main() {
}
