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
	defer file.Close()
	fmt.Println(file, " cause car.txt dont exist till now in this dir")
	file, err = os.Open("cat.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	fmt.Println(file, " shows memory value cause cat.txt exists")
	fmt.Println(file.Name())
	fmt.Println(file.WriteString("awwwwwwwwwwwwwwwwwwwwww"))
	fileRead, err := os.ReadFile("fat.txt")
	if err != nil {
		log.Println("It sems thers no file with name: fat.txt")
	}
	fmt.Println(string(fileRead))
}
