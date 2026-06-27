package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(reflect.TypeOf(args))
	fmt.Println(reflect.TypeOf(args[1]))
}
