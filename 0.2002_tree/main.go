package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirName := filepath.Base(fullPath)
	fmt.Printf("%v\n", dirName)

	a, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, val := range a {
		v, err := val.Info()
		if !v.IsDir() {
			fmt.Printf("└── %s  %d  %s\n", v.Name(), v.Size(), v.Mode())
			if err != nil {
				panic(err)
			}
		}
	}
}

func list() {

}
