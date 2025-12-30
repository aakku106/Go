package concurrency

import "fmt"

// in go select somewhat looks like switch-case (syntex)

const Debug = true

func S1() {
	ch := make(chan int)
	c := make(chan string)

	go func() {
		ch <- 106
	}()
	go func() {
		c <- "weeeeee"
	}()
}

// We already look at this behavior, 1 goutine reaches ch<-106 and waits there infinite and another go rutine
// infinetly waits at c<- ... point and receive never happen, in mean time main dyes and every gorutine was killed, no one noticed
// Lest look at exmple of select(how we catch/receive)
func S2() {
	ch := make(chan int)
	c := make(chan string)

	go func() {
		ch <- 106
	}()
	go func() {
		c <- "weeeeee"
	}()

	select {
	case getch := <-ch:
		if Debug {
			fmt.Println("got value: ", getch)
		}
	}
}
