package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan string), make(chan string)
	go func() {
		for {
			c1 <- "weeeeeeeeeeee"
		}
	}()
	go func() {
		for {
			c2 <- "awwwwwwwwwwwww"
		}
	}()

	for {
		select {
		case message := <-c1:
			fmt.Println(message)
		case message := <-c2:
			fmt.Println(message)
		}
	}

}
