package concurencypattern

import (
	"fmt"
	"time"
)

func Parent() {
	done := make(chan bool)
	go doSomeThing(done)
	time.Sleep(time.Second * 5)
	defer close(done)
}
func doSomeThing(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work from otherFunc")
		}
	}
}

/*
Here what happens is,
In parent we decleared a bool channel, called a func which do somework, where we passed our bool channel

inside that(doSomething) func
we have a infinite loop,
and inside for loop we have select condition
case 1:
when we got some value in done channel, the program it returns and loop stops
But we never passed any value in done, have we ?
Until we get value in done channel, the defult runs


In our Parent func, after calling doSomeThing func, We simulated some work by time.sleep,
It(parent goRutine) does some work till 5 sec, and we close done, and case <-done: runs and doSomething returns, mainitself dies
*/
