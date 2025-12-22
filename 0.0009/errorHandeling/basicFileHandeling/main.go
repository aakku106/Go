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
	// well these seems boring, lest do somethign more intresting in a new cleaner function
	funHandeling()
}

func funHandeling() {
	file, err := os.Create("studentName.txt")
	if err != nil {
		log.Println("Can't creat file named: ", file.Name())
	}
	fmt.Println("A file created with filename: ", file.Name())
	file.WriteString("S.N|---|Name\t\t|\n")
	file.WriteString("1. | Adarasha Gaihre\t|\n")
	file.WriteString("2. | Aakku\t\t|\n")
	file.WriteString("3. | CCN\t\t|\n")
	defer file.Close()
	read, err := os.ReadFile("studentName.txt")
	fmt.Println(string(read))
	// Here 1st we created the file named studentName.txt, then
	// stored its address in file and error in err is there is any, then,
	// checked error with if err!=nil
	// then we write in file using WriteString
	// and used new value read to ReadFile
	// readfile reads file and returns array of bytes(slice of byte to be specific)
}
