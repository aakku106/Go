package main

import "fmt"

func main() {
	fmt.Println(functionInGo("Hellow", "CCN"))
	fmt.Println(functionInGo("Hellow", "CCN"))
	sum, diff, mult := nackedReturnInGo(1, 2, 3, 4)
	fmt.Println("Sum: ", sum, " Diff: ", diff, " Mult: ", mult)
	// lest revied pass by value & pass by refrence
	number := 5
	fmt.Println(multBy2(number))
	fmt.Println(number) // unchanged, cause value was copied into val of called function
	addBy2(&number)
	fmt.Println(number)
	fmt.Println(addBy2(&number)) // called function 2 times so value should be 5+2+2=9
	fmt.Println(number)
	// lest look at array,slice and map in go while passign as arguments
	goArray := [5]int{1, 2, 3, 4, 5}
	goSlice := []int{5, 4, 3, 2, 1}
	goMap := map[int]string{ // domestic animal i love
		1: "cat",
		2: "doc",
	}
	goCointainerTypes(goArray)
	goCointainerTypes(goSlice)
	goCointainerTypes(goMap)
	fmt.Println(arrayTest(goArray))
	fmt.Println(goArray) // whatver we do inside that function it dosent makker in our original array

	sliceTest(goSlice)
	fmt.Println(goSlice)
	fmt.Println(sliceTest2(&goSlice))
	fmt.Println(goSlice)
}

func functionInGo(great, name string) (string, string) {

	return great + " master", name /*
		## function in go
		- inicated with func keywoard
		- paramaters tyoe shall be defined
		- retrn type shall be defined
		- can be multiple return
	*/
}
func nackedReturnInGo(a, b, c, d int) (sum, diff, mult int) {
	sum, diff, mult = a+b, b-c, a*d
	return
	/*
		# Ncked returns in go
		this is nacked return in go
		- it just return the final value of named/variable name given in return value
	*/
}
func multBy2(val int) int {
	return val * 2
	/*
		- this is supper simple go just copied the value of number in val and multyily it by 2 and return that value, original number is not touched
	*/
}
func addBy2(val *int) int {
	*val += 2
	return *val

	/*
					Here we our parameter is val *int, so this function dont takes normal value, it taked the address of value
				so we passed &number as argument in function
				here return *val is just for show since we passed value by refrence and our function not complex that return is just to show how to return the address of val, which has adress of number

		what might be the pros of usign it you may ask ?
		well must simple answe is memory, in pass by value go has to make copy of argument passed (number in our case) which takes space
		not significant amount in our case but it will mater noticable difference if 10K+ users start to call 2 functions
		multBy2 in x time all at once
		and addBy2 in y time all at once
	*/
}
func goCointainerTypes(val any) {
	fmt.Println(val)
	// this function show we cna even assign any in go sounds like typescript ?
}
func arrayTest(val [5]int) (newVal [5]int) {
	newVal[4] = val[4]
	newVal[4] = -5
	val[3] = -4
	return
}
func sliceTest(val []int) {
	fmt.Println(len(val))
	val = append(val, 12)
	val[0] = -1
	fmt.Println(len(val))

}
func sliceTest2(val *[]int) []int {
	*val = append(*val, 12)
	return *val
}
