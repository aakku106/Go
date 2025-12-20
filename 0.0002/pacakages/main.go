package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
	fmt.Println("–––---––––––––––––––––––––Packages in GO–––---––––––––––––––––––––")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())

	// From another file inder same package
	fmt.Println(thisIsFromAnotherFile)
	fmt.Println(thisIsFromAnotherMain)
	fmt.Println(thisIsPrivate)
	fmt.Println(ThisIsPublic)
}

/*
- you may wonder what this is, the topic is package, and i have imported some rendom package/module...You you understand lateron

## Package in Go
- in go a package is the collection of source file (just some/bunch of .go files) in same folder
- in this main.go file i have decleared package as main, so now i cannot make another file in same dir/folder with different package name,
ie.: the name of package in a(same)directory must share a same package name (main in our case)

- What can we do under same package name, it's somwwhat like namespace in other languages
	- in go files under same package name can share variables(any variables) between them(Either Public Or Private)
		- Public variables can be accesed from another package, private cannot be, but we dont have a public and private keywork in go , i  ts lots cleaner than that
		- The variable or Function starting from captial leter are public and smaller are provate.lest look at examples:
Havent you wondered, why desiners of go made P of print statement captial insted of small p ?
its all because fmt.print() would be a private function which couldn't be accesed from outside of that packed or folder in files like this (main.go in 0.0002/packages)

Now you should have notices that in quote.Go the G is captial not small g
But variable i accesed from anotherMain.go has both public and private variables cause its under same package name, share same package name
So

When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter
can be used / accessed.
The recommended style of naming in Go is that identifiers will be named using camelCase,
except for those meant to be accessible across packages which should be PascalCase

*/
