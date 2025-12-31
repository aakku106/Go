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
			fmt.Println(value) // only 1 value ?
			select {
			case ch <- value:
				fmt.Println(value)
				// See we only got our 0th index element:"weee" and not others
			}
		}
	}()
	for i, value := range <-ch {
		fmt.Println(i, ".) Value got: ", value)
	}
}

/*
So here the initial for loop itterated only 1 time,
oki no problem cause we only have 2 goRutines one sends another receives and tham main simpally dyes,
but main point/question or confusion here is whare are those garbage/wired value comming from & and why didnt we
atleast received "weee" ?

also the for loop itterated only once, then from where did even got(receive) for i, value := range <-ch got its value without send
*/
/*
Actually the problem it self is that exact line:
	for i, value := range <-ch
here what we did was not itterated how many time we got data, we actually itterated over the string "weee"
exactly 4 times (Now did you got it whu i used weee and slice of size 4 ?), and those runed we got was
UTF-8 unicode value of "w"-->[119] & "e"-->[101] three time

so our main mistake was <-ch, so lest fix it

*/
//func Fix1() {
//	arr := []string{"weee", "cat", "awww", "lol"}
//	fmt.Println(arr) // we can see its strign
//	ch := make(chan string)
//	go func() {
//		for _, value := range arr {
//			fmt.Println(value) // only 1 value ?
//			select {
//			case ch <- value:
//				fmt.Println(value)
//				// See we only got our 0th index element:"weee" and not others
//			}
//		}
//	}()
//	for i, value := range ch {
//		fmt.Println(i, ".) Value got: ", value)
//	}
//}

/*
 Well we got error here saying:
concurrency/concurencypattern
../concurencypattern/forSelect.go:111:9: range over ch (variable of type chan string) permits only one iteration variable
*/
// THis is where we use close() function

//func Fix1Fix() {
//	arr := []string{"weee", "cat", "awww", "lol"}
//	fmt.Println(arr) // we can see its strign
//	ch := make(chan string)
//	go func() {
//		for _, value := range arr {
//			fmt.Println(value) // only 1 value ?
//			select {
//			case ch <- value:
//				fmt.Println(value)
//				// See we only got our 0th index element:"weee" and not others
//			}
//			close(ch)
//		}
//	}()
//	for i, value := range ch {
//		fmt.Println(i, ".) Value got: ", value)
//	}
//}

// The real culpreata was that i in ch loop it self
// The resion that channel wont give 2 value is cause it only passes the value not store to give index
// un-comment fix1 and fix1fix if you wanna see how it throw error
func Soln() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println("Value got: ", value)
	}
}

// he we got the value we wanted, but how ? lest check

func Test() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
		close(ch)
	}()
	fmt.Println(len(ch))
	for value := range ch {
		fmt.Println("Value got: ", value)
	}
	fmt.Println(ch)
}

// See the length of ch is still 0, althow it passes 4 value
// and ch itself return a address
/*
Actually what heppened here is range ch runs until the ch is closed
Always remember:
- range over channel on receive itterates until it's closed by sender
In above eg, we closed ch after for loop, so it only itterated 4 time, cause there was 4 elements in arr
So, if i had not closed ch after for loop, for value := range ch  would always wait until main itself dies
(But in our case the main won't even die cause that range ch would always block main cause we didnt closed in go func(){}())
*/

// 1 more thing Closing the channel is the sender’s responsibility, we closed ch after looping over arr

func Example() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string, 4)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
		close(ch)
	}()
	fmt.Println(len(ch))
	for value := range ch {
		fmt.Println("Value got: ", value)
	}
	fmt.Println(ch)
}
func Example2() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string, 106)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
		close(ch)
	}()
	fmt.Println(len(ch))
	for value := range ch {
		fmt.Println("Value got: ", value)
	}
	fmt.Println(ch)
}

// See the buffer size really dosen't matter, when we handle when we close channel
// Notice our program is not asynk it's synchironous, we only have 2 thread/goRutine one sends another(main) receive
func Example3() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string, 106)
	go func() {
		for _, value := range arr {
			select {
			case ch <- value:
			}
		}
		close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println("Value got: ", value)
		}
	}()
}

// This func shall never print any value, cause main dies even before anything happens
func Example4() {
	arr := []string{"weee", "cat", "awww", "lol"}
	ch := make(chan string, 106)
	check := make(chan bool)
	go func() {
		for _, value := range arr {
			ch <- value
		}
		close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println("Value got: ", value)
		}
		check <- true
	}()
	fmt.Println(<-check)
}

// Now this example stops until check get any and give any value
// We could also do this insted
func Start() {
	ch := make(chan string)
	value := []string{"cat", "eat", "rat"}
	go Sec(&ch, value)
	go Sec2(&ch)
}

func Sec(ch *chan string, value []string) {
	for _, val := range value {
		*ch <- val
	}
	close(*ch)
}
func Sec2(ch *chan string) (check chan bool) {
	check = make(chan bool)
	for val := range *ch {
		fmt.Println("Got val in Sec2", val)
	}
	check <- true
	defer close(check)
	return
}
