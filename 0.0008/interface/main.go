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

func (r *rectangle) area() float64 { return r.height * r.width }
func (c *circle) area() float64    { return math.Pi * math.Pow(c.radious, 2) }
func (t *traingle) area() float64  { return 0.5 * t.base * t.height }
func (s *square) area() float64    { return math.Pow(s.length, 2) }
