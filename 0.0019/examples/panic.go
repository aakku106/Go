package examples

import (
	"fmt"
	"os"
)

func someWork() {
	defer func() {
		recover()
	}()
	panic("Code isen't Dattebyoo !!!")
	fmt.Println("We cannot reach here")
}

/*
This way the execution of code wont stop and we can handle panic safly
*/
func someBetterWork() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recovered from panic", err)
		}
	}()
	panic("Code isen't Dattebyoo !!!")
	fmt.Println("We cannot reach here")
}

/*
Lest look what's happining here, line by line
defer func till always excurte in last,
than we create panic and passed value(string(panic takes any)) and there is other line of code, which are un-reachable
than defer finally executes, where we recovered panic using builtin func recover()
recover return error (error will be nill when thers no paic)(error will get value<whatver we padded inside painc> and store it in err variable in our case)
and if we get error or value in err ie err != nil, we simpally print somestrings with that err
*/

func SomeFileReading(filename string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	file, err := os.Open(filename)
	if err != nil {
		panic(`File not exist  `)
	}
	defer file.Close()
}
