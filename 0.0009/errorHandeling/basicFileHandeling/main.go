package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Creatig a new file")

	fileAddress, err := os.Create("cat.txt")
	if err != nil {
		errors.New("cant create file")
	}
	fmt.Println(fileAddress)
}
