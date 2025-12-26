package concurrency

import (
	"fmt"
	"time"
)

func C1() {
	fmt.Println("Welcom master.....")
	go get("1")

	// see we got no output at all althow get shall write 1 in terminal we got nothing
	fmt.Println("Just killing time")
	// Even after this we got no output from get
	fmt.Println("Just killing time")
	fmt.Println("Just killing time")
	fmt.Println("Just killing time")
	// even after this much time hte output of get is yet not out
	go get("1")
	go get("1")
	go get("1")
	// even spammin them all wont give anyoutput
	get("1")
	// but that did gave output
	/*
		what happened here is go forked those get function we called in thoir one sebthreads
		its normally called asynchronous behavior
		main func will continue doing its job and won't wait for those job to be done to move on to next job
		ie, whenever we put go infront of any function the parent function won't wait
		that function to do its job and continue doing its job ignoring that function

		In our case, main function just did that it continued doing its job and ignored job of go get(...)
		and when main function dies everything dies, althow we dont have any main function our code,
		the behavior will be same when our program(main thread) dies it kills all other thereds/subthereds/hyperthreads with it
		so go get(...) output were not able to come until the main thread dies,so we saw no putput from go get()
		what can we do to solve this ? well most easy way is to wait for them
	*/

}
func get(s string) {
	fmt.Println("got: ", s)
}
func WaitToGetGet() {
	go get("weeeeee")
	time.Sleep(time.Second)
}
