package main

import "fmt"

type Rectangle struct {
	height, width float32
}

func main() {
	// Method are the special function with ()in front of function name and after func in go
	newRect := Rectangle{
		height: 20,
		width:  30,
	}
	fmt.Println(newRect.area(12.3, 321.1))
	fmt.Println(newRect) // the the original values are unchanged
	fmt.Println(newRect.paramater(12, 321.1312))
	fmt.Println(newRect) // See the original values are changed
}
func (r Rectangle) area(h, w float32) (val float32) {
	r.height = h
	r.width = w
	val = r.height * r.width
	return
}
func (r *Rectangle) paramater(h, w float32) float32 {
	r.height = h
	r.width = w
	return 2 * (r.height + r.width)
}

/*
In the above  example given methods i showed only proves
that
if we pass the value th e original object is untouched, go make a clone of newRect
but if we pass the refrence instead the real value of newRect will change

Simple like in function, but we dont have to pass &newRect this time
*/
