package concurrency

import "fmt"

func T1() {
	channel1 := make(chan string, 2)

	go func() {
		channel1 <- "weeeeeeeeeeeeee"
	}()
	go func() {
		channel1 <- "mmmmmmmmmeeeowwwwwwwwwwwwwww"
		channel1 <- "mmmmmmmmmeeeowwwwwwwwwwwwwww"
		channel1 <- "mmmmmmmmmeeeowwwwwwwwwwwwwww"
	}()
	//	go func() {
	//		channel1 <- "aaaaaaaaaaaaaaaaaaa"
	//	}()

	fmt.Println(channel1)
	mydata := <-channel1
	fmt.Println(mydata)
	mydata = <-channel1
	fmt.Println(mydata)
	mydata = <-channel1
	fmt.Println(mydata)
	mydata = <-channel1
	fmt.Println(mydata)

}
