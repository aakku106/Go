package main

import "fmt"

// You think isen't it too late to learn loop like literally after learning methods and structs& interface
// ans: yes kind of i just forgot

func main() {

	arr := []string{"cat", "dog", "rat", "rat", "bat"}
	for i := range arr {
		fmt.Println(i, " ", arr[i])
	}
	// now now dont yup about what heppend to classical for i:=0;i<len(arr);i++
	// We are not doing that this late, like no way

	// the upper for loop is okiiis, but we actually get 2 values from range as
	for index, key := range arr {
		fmt.Println(" ", index, " ", key)
	}
	// how easy and beautiful it i swe can just avoid arr[i], this is something similar to arr.map((i,v)=>{}) from Js, not be be confused with go map (key value pair of go)

	// also we dont have while loop here, so we simplly to
	i := 0
	for {
		i++
		fmt.Println(i)
		if i == 5 {
			break
		}
		// I here is just for stoping condition cause while loop runs infinetly wntil its base contion gets true
	} //NOTE: it works same in map as array,slice

	// okii, but who wants index I just need value, we can do that oo

	for _, val := range arr {

		fmt.Println(val)
	}
	// simple _ says Go cimpiler to just ignore value
}
