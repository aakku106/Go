package main

func main() {
	var weee rune
	weee = 2
	println(weee)
	weee = '5'
	println(weee)
	weee = '🥺'
	println(weee)
	// its basically int 32
	var sec int32
	sec = '🥺'
	println(sec)
	var cat int8 = 'w'
	println(cat)
	// see go relly dont care itf its text or number, it just represent it ASCII value
	// that's a expected behavior, also see main.c in this same dir, it will give same output
	// Althow text editors will show  C source file not allowed while !using cgo, but both programs will still compile&run
}
