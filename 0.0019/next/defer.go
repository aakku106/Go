package playingWithDefer

import "fmt"

func c1() {
	defer fmt.Println("defer print")
	fmt.Println("normal")
	fmt.Println("normal2")
	fmt.Println("normal3")
}

/*
See how defer is always running in last
*/

func c2() {
	defer fmt.Println("defer print")
	fmt.Println("normal")
	c2Helper()
}
func c2Helper() {
	fmt.Println("normal2")
	fmt.Println("normal3")
}

/*
Even after calling an function defer func always runs at last
this is an intresting behavior,
its like umm, making seperate thread(me be gorutines in go) and running that thread at last after everything ends or may be its like in Js stores somewhere in some shord of queues
*/
func c3() {
	defer fmt.Println("defer print")
	defer fmt.Println("defer print 2")
	fmt.Println("normal")
	c2Helper()
}
func c3Helper() {
	fmt.Println("normal2")
	fmt.Println("normal3")
}

/*
I see the 1st one always executes last may be insted of queue its someshort of stack
*/
func c4() {
	defer fmt.Println("defer print")
	defer fmt.Println("defer print 2")
	fmt.Println("normal")
	c2Helper()
}
func c4Helper() {
	defer fmt.Println("defer print 3")
	fmt.Println("normal2")
}
