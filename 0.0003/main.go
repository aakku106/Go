package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exprementating with time in Go")
	fmt.Println(time.Now())
	fmt.Println(time.Month(1))
	fmt.Println(time.Month(8))
	// time.Month(int) simpally takes number and return string 1=jan,2=fab .....12=dec
	fmt.Println(time.Month(106))

	fmt.Println(time.Weekday(3))
	// same implementation as Month,

	// NOTE:the week start from mondar not sundey

	fmt.Println(time.Hour)
	fmt.Println(time.Hour.Milliseconds())
	fmt.Println(time.Nanosecond.Hours())
	// Hour is the variable having value 1h0m0s ie. one hour zero min 0 sec

	// fmt.Println(ShowTime()) // this is incomplete for now,but will will come in this lateron
	//- these are mostly constants and you are read in the time.go file
	// time.go file has clear comments to make concepts chear
	// anf ablut Hour.int being 3600000000000, cause
	/*
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	*/
	fmt.Println(int(time.Hour))
}

// Function that parse tiem in given function
func ShowTime() (time.Time, error) {
	// the Time you see there is Struck, i have also talked something about Time in comments in buttom
	data := "Sun, 2025/08/27, 1:30 PM"
	layout := "Tus, 2001/03/20, 3:10 AM"
	t, err := time.Parse(layout, data)
	return t, err

}

/*
Time in go is the type which show moment in time
and we also have package named time in go

Note: time is package and Time is type
*/
