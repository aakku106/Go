package main

import (
	"fmt"
	"math"
)

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
	newRect := rectangle{height: 12.32, width: 32.12}
	newCircle := circle{radious: 19.43}
	newTraingle := traingle{height: 42.32, base: 341.4221}
	newSquare := square{length: 1221.32221}

	// Normal just like before
	fmt.Println(
		newCircle.area(),
		newSquare.area(),
		newRect.area(),
		newTraingle.area(),
	) // Well we can this to be consider this as polymorphism

	// We can use interface
	areaCalc(&newRect)
	areaCalc(&newCircle)
	areaCalc(&newSquare)
	areaCalc(&newTraingle)
	// but these are runtime polymorphism
	// dosent this somewhat feels like working with classes, but different
	// the & before struct name is because I used pointer of struct in method defination
}

func (r *rectangle) area() float64 { return r.height * r.width }
func (c *circle) area() float64    { return math.Pi * math.Pow(c.radious, 2) }
func (t *traingle) area() float64  { return 0.5 * t.base * t.height }
func (s *square) area() float64    { return math.Pow(s.length, 2) }
func areaCalc(s Shape) {
	fmt.Println("Area: ", s.area())
}

// There are more examples in example folder
