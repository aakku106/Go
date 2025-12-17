package main

import (
	"fmt"
	"structtest/animal" // Import the animal package!
)

type DomesticAnimal struct {
	Name   string
	height uint16 // lowercase = private (only accessible in this package)
}

// Something like a class but without functions/method inside it

func main() {
	// Using DomesticAnimal from main package
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

	// Using getter/setter
	fmt.Println("Dog height via getter:", Dog.GetHeight())
	Dog.SetHeight(12)
	fmt.Println("Dog height after setter:", Dog.GetHeight())

	fmt.Println("\n--- Now using WildAnimal from animal package! ---")
	// Using WildAnimal from animal package
	Lion := animal.WildAnimal{
		Species: "Lion",
		// Cannot set Weight directly - it's private (lowercase)!
		// Must use setter method instead
	}
	Lion.SetWeight(190) // Use setter for private field
	fmt.Println(Lion)
	fmt.Println(Lion.Species+" weight:", Lion.GetWeight())

	// Can also use functions from animal package
	animal.PrintMonkey()

	// we cna also do
	Dog.Name = "Big dog"
	fmt.Printf(Dog.Name)
	horse := DomesticAnimal{
		Name: "Horse",
	}
	fmt.Println(horse)
	bufflow := DomesticAnimal{
		height: 12,
	}
	fmt.Println(bufflow)
	// See int type returns 0 and string type returns whitespace
}

/*
## Struct in Go
- Sequence or collection of field/variable with its type
- It's something similar to class in OOP languages
- Variable name shall be unique inside a struct
- Struct also follows Go convention of public and private variables

## PUBLIC vs PRIVATE:
- Capitalized names (Name, Species) = PUBLIC = exported = accessible from other packages
- lowercase names (height, weight) = PRIVATE = unexported = only accessible within same package

## IMPORTING RULES:
- You CAN import 'animal' package in 'main'
- You CANNOT import 'main' package in 'animal' (main is special, for executables only)
- You CANNOT have circular imports (A imports B, B imports A)

To access private fields from another package, use getter/setter methods!
*/

// Getter and Setter methods for private field 'height'
func (d *DomesticAnimal) GetHeight() uint16 {
	return d.height
}

func (d *DomesticAnimal) SetHeight(h uint16) {
	d.height = h
}
