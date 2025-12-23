package main

import "fmt"

func main() {
	// Testing fo rerror in function return
	if ok, err := isEven(10); ok {
		fmt.Println("It's even")
	} else {
		fmt.Println(err)
	}
	// Testing if key-value pair exist
	var stock map[string]float64
	s := "Two"
	price := stock[s]
	fmt.Println(price)
	fmt.Println(stock)
	fmt.Println(s)
	if price, ok := stock[s]; ok {
		fmt.Println("Price of ", s, " is ", price)
	} else {
		fmt.Println(s, " not found")
	}

	// stock[s] = 123.321 // we can't do this,so lest use make
	stock = make(map[string]float64)
	stock[s] = 123.321
	fmt.Println(stock)
	stock["Appl"] = 321.123
	fmt.Println(stock)

	if price, ok := stock["appl"]; ok {
		fmt.Println("Appl stock value: ", price)
	} else {
		fmt.Println("aal not found")
	} //cause a is capital A
	if price, ok := stock["Appl"]; ok {
		fmt.Println("Appl stock value: ", price)
	} else {
		fmt.Println("aal not found")
	}
	// this is simple value checking nothing special althow it look verbose its very simple value checking

	// lest play with interface (Type assertion they say)
	// thers seperate folder in 0.0009/assertion
	var i interface{} = "awwwwwwwwwwwwww"
	if n, ok := i.(int); ok {
		fmt.Println(n, " is int")
	} else {
		fmt.Println("Its not int")
		if n, ok := i.(string); ok {
			fmt.Println(n, " is String")
		} else {
			fmt.Println("Its not string")
		}
		//this is basically like comparing value or key in map, but we need to know to use (int),(string), or out intended/needed type

		// Here you may think that variable n is just useless we could just use i insted, yes you can always do that
		if _, ok := i.(string); ok {
			fmt.Println(i, " is String")
		} else {
			fmt.Println("Its not string")
		}
	}
}

func isEven(n int) (bool, error) {
	if n&1 == 1 {
		return false, fmt.Errorf("It's Odd")
	}
	return true, nil
}
