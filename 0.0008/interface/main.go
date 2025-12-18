package main

import "math"

type Shape interface {
	area() float64
}

type rectangle struct {
	height, width float64
}
type circle struct {
	radious float64
}
type traingle struct {
	height, base float64
}
type square struct {
	length float64
}

func main() {

}
