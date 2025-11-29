package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
	fmt.Println("–––---––––––––––––––––––––Packages in GO–––---––––––––––––––––––––")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
}
