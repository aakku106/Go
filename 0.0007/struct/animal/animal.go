package animal

import "fmt"

// WildAnimal is an exported struct (starts with capital letter)
// This can be used in other packages that import 'animal'
type WildAnimal struct {
	Species string  // Exported field (accessible from other packages)
	weight  float64 // Unexported field (private, only accessible in animal package)
}

// Getter for private field 'weight'
func (w *WildAnimal) GetWeight() float64 {
	return w.weight
}

// Setter for private field 'weight'
func (w *WildAnimal) SetWeight(newWeight float64) {
	w.weight = newWeight
}

// PrintMonkey is an exported function (starts with capital letter)
func PrintMonkey() {
	fmt.Println("Monkey from animal package!")
}

/*
KEY POINTS:
1. This package CANNOT import main (main is special)
2. This package defines its own types (WildAnimal)
3. Other packages can import THIS package and use exported items
4. Exported items start with CAPITAL letters (WildAnimal, Species, GetWeight)
5. Unexported items start with lowercase (weight)
- and special rule, import cannot be circular
## what does circular import of packet means
eg: we have package animal and another package build in build/build here
- we can import animal in build
- can import built in animal
- but cannot import built in animal and animal in built at one time
ie. Either import built in animal oe import animal in built but not each other at same time
*/
