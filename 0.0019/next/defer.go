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
	c3Helper()
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
	c4Helper()
}
func c4Helper() {
	defer fmt.Println("defer print 3")
	fmt.Println("normal2")
}

/*
Yes this do act like stack
*/
func d1() {
	defer fmt.Println("Exiting D1")
	defer d1Point1()
	defer d1Point2()
	defer d1Point3()
	defer fmt.Println("Starting D1")
}
func d1Point1() {
	fmt.Println("Printing at last")
}
func d1Point2() {
	fmt.Println("Printing at 2ndlast/2nd")
}
func d1Point3() {
	fmt.Println("Printing at 1st")
}

/*
As aspected, this show a consistent stack behavior everythere,
thsi is perfect to close servers or actually db connections even before we start connection it will insure db will close at last
This is liek destructor in C++ ~
*/
// Next we will look at some examples
// Navigate to ../examples/files.go
