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

the truth is T1 is the example of receiver waits until send happens,
but i just told it was example of send waits until receive happens, yes i did told that, & I am correct
, Becaue you never know if goRutine will get ch <- 10  inside go func(){}() first or
	x := <-ch inside func T1(){}

// So if goRutine reach ch <- 10  first than its example of send waits until receive happens, and
if 	goRutine reaches x:=<-ch its example of Receive waits until Send happens
*/
// Now lest look at another concept, that channels dont hold the data, they are just pipe passing down the data

func T3() {
	ch := make(chan string)
	// Here we created a pipe name ch which can carry(only) string type
	fmt.Println("length of ch: ", len(ch))
	go func() {
		ch <- "Cat"
		fmt.Println("length of ch: ", len(ch))
	}()
	fmt.Println("length of ch: ", len(ch))
	x := <-ch
	fmt.Println("length of ch: ", len(ch))
	fmt.Println("Length of X: ", len(x))
	fmt.Println(x)
}

/*
See the length of ch is always 0
no matter where we check, because its not holding value like array or slice, it only passes value when receiver and sender shacks the hand/ meets

and length of x is 3 because length of string means how many caracter does it contain
*/

// That was un-buffered channel in go, now lest look at buffered channel in go in function B1
func B1() {
	ch := make(chan int, 2)
	go func() {
		ch <- '?'
	}()
	x := <-ch
	fmt.Println(x)
}

/*
In func B1 we created a pipe named ch which can only transfer 2 int at a time
it's basically like buffered, but its max capacity is 2 unlike un-buffered max cap was 1

in above example func B1
we are only sending 1 value and receiving 1 value we sent
but you might think why we sending ?, wasent it supposed to be int type, than you shall check /Go/0.0010/runes 1st

And now next qtn may be the capacity was 2, but we only sent 1 value and it worked, yes
here 2 is max cap not min cap
In go Channels min cap is always 1
// Lest look at another example B2
*/
func B2() {
	ch := make(chan string, 2)
	go func() {
		ch <- "106"
	}()
	go func() {
		ch <- "weeeeeeeeeeeeeeee"
	}()
	x := <-ch
	fmt.Println(x)
}
func B3() {
	ch := make(chan string, 2)
	go func() {
		ch <- "106"
		ch <- "cattt"
	}()
	x := <-ch
	fmt.Println(x)
}

/*
here in B2 we got valuw weeee which was assigned after 106,
but in B3 we got value 106 which was assigned before cattt, and why is that ?
*/
func B4() {
	fmt.Println("B4 starts here...............")
	ch := make(chan string, 2)
	go func() {
		ch <- "106"
	}()
	go func() {
		ch <- "weeeeeeeeeeeeeeee"
	}()
	x := <-ch
	y := <-ch
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("B4 Ends here...............")
}

/*
Her we got weeeeee.... before and 106 after why ?
*/
func B5() {
	fmt.Println()
	fmt.Println()
	fmt.Println("B5 starts here...............")
	ch := make(chan string, 2)
	go func() {
		ch <- "106"
		ch <- "cattt"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("B5 Ends here...............")
	fmt.Println()
}

/*
Here we got value as aspected, 1st 10y than cattt,
(
	assigning value to variable wont change anything its same  if you do
	x := <-ch; fmt.Println(x)  or
	fmt.Println(<-ch)
)

In go
one law you must accept
A channel preserves order per sender, not across senders

This single sentence explains everything you’re seeing.
- FIFO within one goroutine
- No ordering guarantee between goroutines
- Buffered channels only change blocking, not ordering guarantees
*/

/*
So in case of B2 or B4
-	Two goroutines are launched
-	Both attempt to send
-	Channel has buffer size 2 → both sends succeed immediately
-	Which goroutine runs first is undefined

So this means in go
NOTE: The order in source code does not matter once goroutines exist.

So the buffer ends up like
either ["106", "weee"] or ["weee", "106"]
our output reflects arrival order, not source order
*/

/*
AND in case of B3 or B5
Why this prints 106 (always)

Because now there is one sender

Inside a single goroutine:
- Instructions are executed in order
So the channel buffer becomes(Always):["106", "cattt"]
FIFO is respected because there is only one sender, this is clean and predictable case
*/
// NOTE: In your computer B2&B4 may completly fip, but when i am trying its always weee.. 1st and 106 2nd (This may be in your case/computer/envirenment)
// Who knows which GoRutine will run first
// This is runtime-level questions.

// Now let us see example where we have more than 2 value in channel with 2 buffer size in B6

func B6() {
	fmt.Println()
	fmt.Println()
	fmt.Println("B6 starts here...............")

	ch := make(chan int, 2)
	ch <- 12
	ch <- 14
	ch <- 16
	fmt.Println(<-ch)

	fmt.Println("B6 Ends here...............")
	fmt.Println()
} // This function will be always be block no matter what
func B7() {
	fmt.Println()
	fmt.Println()
	fmt.Println("B7 starts here...............")

	ch := make(chan int, 2)
	go func() {
		ch <- 12
		ch <- 14
		ch <- 16
	}()

	fmt.Println(<-ch)

	fmt.Println("B7 Ends here...............")
	fmt.Println()
} // Bu this function will, its simplly because the ch<- value in B6 blocks the only one gorutine that was operating on that function
// but since we created go func() which creates one more goruting, so one receives value and another GoRutine sends the value

func B8() {
	fmt.Println()
	fmt.Println()
	fmt.Println("B8 starts here...............")

	ch := make(chan int, 2)
	go func() {
		ch <- 12
		ch <- 14
		ch <- 16
	}()

	fmt.Println("B8 Ends here...............")
	fmt.Println()
} // Even this will run in buffered-channel

func B9() {
	fmt.Println()
	fmt.Println()
	fmt.Println("B9 starts here...............")

	ch := make(chan int, 2)
	go func() {
		ch <- 12
		ch <- 14
		ch <- 16
	}()

	fmt.Println("B9 Ends here...............")
	fmt.Println()
}
