package main

import "fmt"

var (
	prorityValue int
)

func prorityQueue() {
	fmt.Println("Accessing Protity Queue.....")
	fmt.Println("Protuty Queue Accesed")

}

/*
Lest see how we make  prority queue:
1. 1st it shall follow fifo rele: first in 1st out,which means i have to 1st make a normal linear queue(not circular queue cause i have a idea to use slice like a ummm some shord of garbage collector thing like, when lest say slice is larger than 50, ie. the front is larger than 50 than only do [front:], and front=0; i think it should be ~ O(1)may be lest cee)
2. i need a prority, so lest decleare a var called prority, where we can give values from 0 to 4, lower the value higher the prority, I think this is good idea for now
*/
