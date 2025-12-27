package concurrency

import "fmt"

/*
## Channels in Go
- simpally talking channel are like a pipe
	put somethign in 1 side, and get that thing out from other side

### RULE
- Send wait untill Receive happens,
- Receive waits until Send happens
? sounds confusing, you are in right path, continue reading
*/

// There are generally 2 types oc channels in Go
//	i.	Un-Buffered channel
//	ii. Buffered Channel

// Lest look at un-Buffereed channel 1st in fucn T1
func T1() {
	ch := make(chan int)
	// Here we created a pipe name ch which can carry(only) int type
	go func() { ch <- 10 }()
	x := <-ch
	fmt.Println(x)
	// you might think what are those wired symboles and wired code writtem, than continue reading
}

/*
In above function T1, we created a pipe named ch which can only carry int,
than go func()starts,
and reaches: ch<-10 and
waits there/ STOPs there, it won't continue, its waiting
Why? : Because no one is receiving yet
and in main, main reaches x:=<-ch
now sender and receiver meet, value(10) is passed across through ch pipe, and both continue
This setisfies 1sr rule:
- Send wait untill Receive happens,
, but what about receive waits until send happens,
// Lest look at T2 function (Also an example of un-buffered channel)
*/

func T2() {
	ch := make(chan string)
	// we created a pipe named ch which can only carry sting
	x := <-ch
	//	fmt.Println(x)
	go func() { ch <- "weee" }()
	fmt.Println(x)
}

/*
Now this will never work, because x:=<-ch blockes the thread(gorutine/green thread)
So untill and unless go sees go func(){...} it will never create a new go rutine which would go inside go func(){}() and receives the value,
than how do even we see example of
- Receive waits until Send happens
*/
