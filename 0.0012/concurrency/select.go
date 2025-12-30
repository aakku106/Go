package concurrency

import "fmt"

/*
## we know
	1.	A send or receive may block
	2.	Multiple goroutines may race to talk to you
	3.	The program can end while goroutines are stuck

So the real question becomes:

“What if I’m waiting on multiple channels, but I don’t know which one will be ready first?”

That is precisely what select solves.
*/

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

// this is like a switch case statement
// now the main waits until it gets something in getch, but what about c ?
func S3() {
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
	case getc := <-c:
		if Debug {
			fmt.Println("got value: ", getc)
		}
	}
}

// se we got both value here, but that's wasent aspected
// a select blocks until one of its cases runs(same like switch case)

/*
	select blocks until one case can run
	if multiple cases are ready, one is chosen randomly
*/
// than why did we got both value here ?

/*
actually we are only getting 1 value from func S3 that 106 was from func S2
*/

/*
But we have a very serious problem in our code here, the goRutines are piling up
In s1 2 go rutines are just waiting forever
In s2 1 is waitign for ever
and in s3 1 is blocked for ever,

Well its big deal here, cause our main dyes fast but in real case/production a main almost never dies.

This is why production Go code:
	-	closes channels
	-	uses WaitGroup
	-	uses context.Context
	-	or structures ownership clearly

select without lifecycle management is dangerous
*/
//NOTE: select waits for one channel operation, executes exactly one case, then stops — everything else is your responsibility

// Next we will look into concurrency pattern (./concurencyPattern/*)
