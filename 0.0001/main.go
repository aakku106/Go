package main

import "fmt"

func main() {
	fmt.Println("i like Cat")
	a := 1
	b := 2
	c := a + b

	fmt.Println(a, b, c)
	cat := "a domestic animal"
	bol := true
	rat := "eat it"
	fmt.Println(cat, rat, bol)

	// array

	arr := [5]int{1, 2, 3, 4, 5}
	anotherArray := [2]string{"cat", "rat"}
	fmt.Println(arr, anotherArray, anotherArray[0], anotherArray[1], arr[4])

	// Slice
	// thay are like brothers ot array, but more opean minded xd..

	sli := []int{1, 2, 3}
	fmt.Println(sli)
	sli = append(sli, 4)
	fmt.Println(sli)
	sli = append(sli, 4)
	fmt.Println(sli)
	NewSlice := sli[1:3]
	fmt.Println(NewSlice)
	// this is somewhat familiar to Python ...

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)
	ss := s[:0]
	printSlice(ss)
	sss := s[1:]
	printSlice(sss)
	// of curse we can reuse a variable
	sss = sss[2:]
	printSlice(sss)
	sss = sss[:2]
	printSlice(sss)
	sss = sss[1:2]
	printSlice(sss)
	// notice the capacity

	/*

									- slice is the pointer to underlying array, slice dont own arraym it just points

									- In above code, with s we decleeared a slice or some array with capacity, length 10

								- capacity means how much array can hold, and leangh means how much its currently holding.

								- in ss, we copited s,but [:0] means copy item/element less than 0, so its length became 0, but since ss is pointing to s in memory somewhere, its capacity is still 10 (re-view code aand watch results to graps whats i am saying)

								- now here in sss, we pointed to s, but copied elements after 0th element, so length becomes 9 and since we copied from 1st element capacity also becomes 9

									- Now after that we just re-used same slice/array(pointed in memory) sss after that we did sss= sss[2:], ie copy sss in sss, from index 2(0,1,2...not 1,2...ie. index 2 = 3rd element of array )
						 so the capacity of sss now changes from 9 to 9-2= 7;
					- now did you got confused here, in slice [2:] ie 3rd element, but 9-2, shouldn't it be 9-3 ?
					- Confused if yes lest review our concepts again,
					1. Slice do not own array, it is just pointer to underlying array in memmory.
			2. New capacity = old capacity - startingIndex

		lest look at another string example to make concepts clear.


	*/

	string := []string{"My", "Name", "is", "ccn", "okii"}

	stringPrint(string)
	// lest slice them
	stringMy := string[0:1]
	stringPrint(stringMy) // its capacity is 5 no problem and length is simplly 1 cause oe only have 1st index or 0th element
	newSlice := string[1:]
	stringPrint(newSlice) // Here length is 4 and capacity is 4
	// lest look why is length 4, cause simplly, if we start from index 1 or 2nd element, we only have left Name, is, ccn, okki
	// and lest move to intresting thing, why is capacity 4,
	/*
		Here,
		newSlice points to string (original array with capacity 5 & length also 5) ["my","Name","is","ccn","okki"]
		and says my startign index will be 1st index of string or 2nd member of array string.
		which makes capacity of newSlice 4 and new array is [name","is","ccn","okii"]
	*/
	// I think the above comment clearly explain behavior of slice, we will be talking on why slice behaves like this latter on
	// also i will make a seperate file for all example and explains of slice seperatly.
	// Now lest go to next step we will make slice of slice
	newSlice = newSlice[1:]
	// If you think its output will be same as before, than you are wrong.
	stringPrint(newSlice)
	// the output is [is ccn okii], which is slice of a slice, which means,
	// initially newSlice sayed my 0th index will be 1st index of string ie. ["<this is 0th index of new slice>Name","is","ccn","okki"] ["<0th index of string>My", "<0th index of new slice>Name", "is", "ccn", "okii"]

	/*
				and now ["My", "Name", "<0th index of new Newstring>is", "ccn", "okii"], which is  strings[2:] or [1:] of newString
		["name","is","ccn","okii"] i.e. new newString is ["is","ccn","okii"]
	*/
	// Do you wonder why i sayed new newString is [2:] of string, does that makes you confuse, good you ar learnign than.
	// One more critial never name your variable string.
	/*
												# lest try to solve new newSlice is [2:] of original string slice, we can say so because

												A slice is stored like,
												- pointer → original array (at some offset)
										- length
										- capacity

								Every time you slice:
									•	you MOVE the pointer forward
									•	you REDUCE capacity
									•	you CHANGE length



					a:= [x:]
					b:=a[y:], then
					b == arr[x+y:]

		hope this dosent make you more confused, read this 2 to 3 times, take time, look at code, look output and read comments again , try doing aaa:=string[2:] you shough be gettign concept now
	*/

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
func stringPrint(s []string) {
	fmt.Printf("\nNew Array:\t%v\t Length\t%d\tCapacity\t%d\n", s, len(s), cap(s))
}

// everything written in this file (time: 2025,nov 27, 10:46 PM) are not offically conformed, this is just byproduct of my brainstroming of Go's slice based on its behavior, its my 1st say of learning GO,
// toold used: alacrity(terminal emulator), Nvim (code editor),offical go doc (may be that was official i just rad data dype in go and started exploring my self) and used chatGPT 5 to verify this file (It told everything in correct in this file except grammer, but i still dotn belife that stupid bot, how could even this all be correct in 1st try of go)
// you still intrested, i tryed to prove this in proveSlice.go file and all slice content of this file in in seperate go file named slice.go (may be in different dir as this file<main.go>)
