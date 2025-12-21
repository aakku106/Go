package main

import "fmt"

type cat interface{}

// this is how any keywoard is made, actually any is not even keywoard
// its interface with no methods
//
//	In Befores example given for calculation of area of different sizes, i used (name *shape name), but we dont always have to do that in interface
type shape interface {
	set(...float64)
	area() float64
}
type rectangle struct{ length, breadth float64 }

func printArea(s shape) { fmt.Println("Area: ", s.area()) }

func (r *rectangle) set(nums ...float64) {
	switch {
	case len(nums) == 2:
		r.length = nums[0]
		r.breadth = nums[1]
	default:
		fmt.Println("Not enough values")
	}
}
func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func main() {
	newRect := rectangle{}
	go newRect.set(123.321, 321.123)
	fmt.Println(newRect.area()) // Here i printed area in main function, i can insted use iinterface
	printArea(&newRect)
	printArea(&rectangle{}) // We can even do this but it will provly print area:0
	// Playin gwith empty interface
	emptyInterface()
}
func emptyInterface() {
	// lest do some fun with empty interface
	// they can be powerful
	var number cat
	number = 5
	fmt.Println(number)
	number = "cat"
	fmt.Println(number)
	/*
				Wasent it fun ? any is just same
				Go: here wo do strong type safty at compile time to avoid wired behavior of code
		meanwhile users: Lest turn go into Js by using any everywhere
	*/
	domesticAnimals := map[cat]cat{
		1:     "cat",
		"dog": 2,
		// We breaked every rule of Go
		// May be useful in JSON
	}
	printEmtyInterfacetypes(domesticAnimals)
	// lest test array/slice
	var something cat
	something = []cat{1, 2, "cat"}
	printEmtyInterfacetypes(something)
	// we can also do
	something = map[any]cat{
		"cat": 1,
		2:     "bat",
		true:  false,
		// why i used any here, why not ?xd
	}
	printEmtyInterfacetypes(something)
	// Lest look at more fun things
	emptyInterfaceVariadicType(fmt.Println())
	emptyInterfaceVariadicType(fmt.Println("buhuhahaha"))
	emptyInterfaceVariadicType(1, 2, 3)
	emptyInterfaceVariadicType(fmt.Println("cat"))
	emptyInterfaceVariadicType(fmt.Print("cat"))
	emptyInterfaceVariadicType(fmt.Printf("cat"))
	emptyInterfaceVariadicType(fmt.Printf(""))
	emptyInterfaceVariadicType(fmt.Print(""))
	emptyInterfaceVariadicType(fmt.Print())
	// basicallly ...cat makes a array of cat type, simple

}
func printEmtyInterfacetypes(v cat) {
	fmt.Println(v)
}
func emptyInterfaceVariadicType(v ...cat) {
	fmt.Println(v)
	//Use empty interface only when the structure is un-known
	// Its compiles, dosent mean its correct
	// IT cant be concurent
}
