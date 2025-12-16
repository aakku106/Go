package main

import "fmt"

func main() {
	thisIsMap := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(thisIsMap)
	for i := range thisIsMap {
		fmt.Printf("\t%d\n", thisIsMap[i])
	}
	stringMap := map[int]string{
		1: "cat",
		2: "dog",
		3: "rat",
	}
	fmt.Println(stringMap)

	//we can also do
	stringMap[4] = "Monkey"
	fmt.Println(stringMap)
	// i did told that we have to give unique key in a map, but what of we give same key ?
	stringMap[4] = "fly"
	fmt.Println(stringMap)
	// the answer is that the value in key 4 chances or to be specific the value of key updates
	// what if we wanna pop the value from map (remove value from map)
	delete(stringMap, 4)
	fmt.Println(stringMap)
	// simplly use delete function
	// delete is defined in builtin.go between copy and len

	// Lest go to more confusing part of map now

	fmt.Println(thisIsMap[4])
	// it returns 0 but we dont have 0 in our map in thisIsMap
	// this can be confusing we dont even have 4 key in our map and go returns a wired number, but if we do same with stringMap
	// it will return nothing or rather whitespace
	// so to check if key exist or not we can do
	value, exist := thisIsMap[4]
	fmt.Println("value:  ", value, "  exist:  ", exist)

	value, exist = thisIsMap[1]
	fmt.Println("value:  ", value, "  exist:  ", exist)

	// See the it return 2 value a boolean and value of key, so if key dont exist it returns bool(false) and (0 or whitespace) and if exist it returns value of key and bool(true)

}

/*
## map in go
- in go map is the built in data-type to store key value pairs just like object in js or dictonaries in python, but there are some defference

- in go its more explicit than in js
- we have to specisy the type of key and type of value
	in thisIsMap both key & value
- we can transverse map just like array with index here thisIsMap has length 2
In second example i have given key as int type and value as string type

NOTE: only one thign that is wired is that you have to give comma ',' even in last

NOTE: the key should be unique

NOTE: If you got confused in anything in this fil ejust go to its defination, it will make dought clear.

*/
