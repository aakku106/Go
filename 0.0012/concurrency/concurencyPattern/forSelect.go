package concurencypattern

import "fmt"

func Eg1() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	for _, value := range arr {
		ch <- value
	}

}

// you can feed the chan like this, but this func will be always blocked
// gorutine will alwals be blocked in 1st itteration of for loop in ch<-value, cause we nevered receive it
func Eg2() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	for _, value := range arr {
		select {
		case ch <- value:
		}
	}
}

// this will be also always clocked, cause we never received value, but this time we used forSelect pattern
// ie. whenever ch gets value our flow will move forward
func Eg3() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	for _, value := range arr {
		select {
		case ch <- value:
		}
	}
	for value := range <-ch {
		fmt.Println("Value got: ", value)
	}
}

// Well we did receive the value, but our one and only gorutine was always stopped at 1st itteration of our 1st for loop
// Lest solve that
func Solve() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
	}()
	for i, value := range <-ch {
		fmt.Println(i, ".) Value got: ", value)
	}
}

// Well now we did get value as aspected, but we got runes(int32) insted of our original strings
// Les't debug why,where that heppened line by line
func Solve2() {
	arr := []string{"weee", "cat", "awww", "lol"}
	fmt.Println(arr) // we can see its strign
	ch := make(chan string)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
				fmt.Println(value)
				// See we only got our 0th index element:"weee" andd not others
			}
		}
	}()
	for i, value := range <-ch {
		fmt.Println(i, ".) Value got: ", value)
	}
}
