package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Creatig a new file")

	fileAddress, err := os.Create("cat.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(fileAddress)
	file, err := os.Open("car.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file, " cause car.txt dont exist till now in this dir")
	file, err = os.Open("cat.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(file, " shows memory value cause cat.txt exists")
	fmt.Println(file.Name())
}
