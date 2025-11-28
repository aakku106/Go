package main

import "fmt"

func proveSlice() {
	// What i gonna show here ?
	// - Slice do not own the array, it is just pointer to underlying array in memmory

	fmt.Print("\033[H\033[2J")
	fmt.Println("-----------------Prove slice-----------------")
	arr := []string{} //capacity 0 simple
	printSlic(arr)
	arr = append(arr, "Cats") //capacity 1 simple
	printSlic(arr)
	arr = append(arr, "Eats") //capacity 2 simplee
	printSlic(arr)
	arr = append(arr, "Rats") //Capacity 4 wtf
	printSlic(arr)
	// lest do one more time than you will understand the pattern
	arr = append(arr, "and") // see, now capacity is still 4 and length is also 4.
	printSlic(arr)
	arr = append(arr, "Becomes") // Here capacity becomes 8 and lenght is 5
	printSlic(arr)
	arr = append(arr, "Fat") // Capacity is still 8 and length gets 6
	printSlic(arr)
	// so, Go simplly doubles the size of array or slice when its capacity gets filled,
	// Again slice points array with fixed capacity, go just doubles the size of pointed array when capacity is full
	// When capicity ie slice fills up and we append more elements, Go does: old capacity x 2 (oldCapacity*2)
	newArr := arr[:]
	printSlic(newArr)
	newArr[5] = "Super Fat"
	printSlic(newArr)
	printSlic(arr)
	// See, what we did was just changed 6th element or 5th index of slice newArr,
	// which chaged the value of newarr to [Cats Eats Rats and Becomes Super Fat] from  [Cats Eats Rats and Becomes Fat]
	// But also 6tm element of arr also changed to [Cats Eats Rats and Becomes Super Fat] from [Cats Eats Rats and Becomes Fat]
	// This why i told slice of slice [1:]of [1:] is [2:] of slice string

	// One more thong to know about slice, althow the capacity of new array is 8, we cant do
	newArr[7] = "last index"
	// we will get runtime error like:
	//panic: runtime error: index out of range [7] with length 6
	// Do you know why is this ?
	// I herd it makes go safe
	// LEst explore this too
	/*

				In Go:
					Length (len) = how many elements your slice currently has.
					Capacity (cap) = how many elements the underlying array can possibly hold before resizing.
		- Slices only allow access from index 0 to len-1.
					Even if capacity is bigger, you can’t use those extra slots directly.

				those extra slots are not part of the slice’s length.
				They exist in memory, but the slice doesn’t “own” them in its metadata.
	*/
	// simplly this means we cannot access elements out of length, so we can say our length is the size of slice,
	// and capacity indicates max length of that array(array not slice because slice simplly points to array in memory adn arrays has fix size, capacity denotes the size of that array, and length denotes size of slice you using. So you can't change the value of underlying array until length of slice reach to that index).

	// Capacity shows potential.
	// Length shows what’s currently real.
	// You can NEVER index beyond length.
	// You CAN append up to capacity before Go reallocates.

	// This much of depth is oki for our biginner level.

	// Everything i explained in this folder is only applicable with slice in go, if you use same logic for array (what i explained about length and capacity) will be wrong, array are simple just like in C(at this lvl you dont need to knwo more,you will learn with time.).

}
func printSlic(s []string) {
	fmt.Printf("\ngiven Array:\t%v\tLength:\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}

// This program ends with an runtime error,Try solving this error.
// Next go to error/main.go file, we will learn error handeling in Go in this same code base
