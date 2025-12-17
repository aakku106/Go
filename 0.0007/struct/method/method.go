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

/*
NOTE: method don’t let you do something you can’t do with functions.
They let you do things cleaner, safer, and more scalable.

than you might why bother to use it ?
it will make all scence while using interfaces

also looking at
newRect.area(12, 30)
&
area(newRect, 12, 30)
the upper one is more cleaner and explicetly saying what we are doing lest look at another similar eg
normal Function
 area(rect)
 perimeter(rect)
 scale(rect)
Methods
 rect.area()
 rect.perimeter()
 rect.scale()
// In function area what ? scale what ?, in methods o rect.ares, react.scale or traingle.perimeter

methods are not for performance or functions lacks or because of pointer only
what it is for
 Behavior ownership
 API clarity
 Interface compatibility
 Large-code sanity
 Expressing intent
*/

/*
Another qtn you may ask is, why go avoided classes,but used methods

Why no classes?

Classes usually bring:
		inheritance chains
		fragile base classes
		tight coupling
		“is-a” abuse
		override hell

Go’s designers (Rob Pike, Ken Thompson, etc.) hated this.

They saw:

“Most real systems use composition, not inheritance.”

So Go said:
		 No classes
		 No inheritance
		 No virtual methods
		 No constructors
*/
