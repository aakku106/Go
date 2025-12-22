package main

func main() {
	a := 12
	b := 106
	println(a)
	println(b)
	if true {
		a = 106
		b := 12
		println(a, " a")
		println(b, " b")
	}
	println(a, " a") // here value of a changees simple.
	println(b, " b") // here the value of b returns to 106

	// Is's good idea not to use variable showind(re declearing b inside other func), these creates confusion and
	// in go its even more diffucult to differicante betn a:=10 and a=10,':' hard to notice
}
