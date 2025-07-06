package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("weeeeeee\n")
	num := 14
	fmt.Println(num + 1)
	print(`we`)
	cat := `cat are mew mew`
	va1, var2, var3 := 1, 2, 3
	println(cat, va1, var2, var3)
	const fatCat = `meowwwwwwwwwwww`
	function()
	function2(`cat hehe`)
	fmt.Println(add(1, 2))

	fmt.Println("Before asynk function")

	go asynk(1) // main dosent wait for it so we wont even see its output at all
	go asynk(2)
	go asynk(3) // but if want to see output we can make some delay after asynk function call

	time.Sleep(time.Second)
	fmt.Println("After asynk function")
}
func function() {
	println(`weeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee!!!!!!!!!`)
}
func function2(cat string) {
	fmt.Println(cat)
}
func add(a, b int) int {
	return a + b
}
func asynk(num int) {
	fmt.Println(num)

}
