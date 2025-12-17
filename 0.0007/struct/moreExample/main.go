package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length, breadth float32
}
type circle struct {
	radious float64
}
type traingle struct {
	base, height float32
}

func main() {

	square := rectangle{
		length: 21.63, breadth: 21.63,
	}
	newCircle := circle{
		radious: 21.23,
	}
	rtTraingl := traingle{
		base: 15.2, height: 19.8,
	}
	fmt.Println("Area of Square: ", square.area())
	fmt.Println("Area of Circle: ", newCircle.area())
	fmt.Println("Area of Traingle: ", rtTraingl.area())

}

func (r *rectangle) area() float32 {
	return r.length * r.breadth
}
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radious, 2)
}
func (t *traingle) area() float32 {
	return (t.height * t.base * 1 / 2)
}
