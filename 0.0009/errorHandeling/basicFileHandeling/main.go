package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Creatig a new file")

	fileAddress, err := os.Create("cat.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileAddress)
	file, err := os.Open("car.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file, " cause car.txt dont exist till now in this dir")
	file, err = os.Open("cat.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file, " shows memory value cause cat.txt exists")

}
