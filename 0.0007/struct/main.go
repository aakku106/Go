package main

import "fmt"

type DomesticAnimal struct {
	Name   string
	height uint16
}

// Something like a class but withour functions/method inside it

func main() {
	Dog := DomesticAnimal{
		Name:   "Dog",
		height: 10,
	}
	fmt.Println(Dog)
	fmt.Println(Dog.Name+" height:", Dog.height)

	Cat := DomesticAnimal{
		Name:   "Cat",
		height: 3,
	}
	fmt.Println(Cat.Name+" height:", Cat.height)
	// and this is something similar to making object from clas called instances of struct in Go
}

/*
## Struct in Go
- Secuence or collection of field/variable with it type
- its Something similar to class in OOPS languages
- var name shall be unique inside a struct
- struct also follows Go convention of public and private variables

So when we are calling something.height in package main its accesable, but from other packages lest say animals/animals.go package animal,
we cannot access something.height, but something.Name is accesable
this can create confusion in 1st but you will be get used to it after some practice
now you wonder, what if we need value in package animal and we need to give value from animal package
Anmswer is simple,
just add getter and setter function in package main ie. this file
*/

func (d *DomesticAnimal) GetHeight() uint16 {
	return d.height
}
func (d *DomesticAnimal) SetHeight(h uint16) {
	d.height = h
	// Now look at animal/animal.go package animal
}
