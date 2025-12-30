package test

import (
	"concurrency"
	"concurrency/concurencypattern"
	"testing"
)

func testGoRutines(t *testing.T) {
	// concurrency.WaitToGetGet()
	concurrency.C1()
	// we won't see any output from go get function of C1, function (but those get function will be printed because we waited for some seconds in WaitToGetGet functer below, thats the resion all go get from c1 wont be printed cause timwA we waited weren't enougth to print them all)
	concurrency.WaitToGetGet()
	// if you call  concurrency.WaitToGetGet() before concurrency.C1() you can observe only wee will be printed

	// that's basic demosestration of go rutines
}

// The small t instead of T is used here caus ei didn't wanted the output of goRutines file while i was testing other files
// If you want to test gorutines.go file just change from t to T

func testChannels(t *testing.T) {
	concurrency.T1()
	//	concurrency.T2()
	//
	// why i commented this t2 function is becaue will will block the main gorutine and waits till inifinit until someones kill the process
	// try Uncommenting if you dont get what am i trying to explain
	concurrency.T3()

	// Now test for buffered channel

	concurrency.B1()
	// just 1 value assigned till here
	concurrency.B2()
	concurrency.B3()
	concurrency.B4()
	concurrency.B5()
	// concurrency.B6()
	concurrency.B7()
	concurrency.B8()
	concurrency.B9()

}
func testGoSelect(t *testing.T) {
	concurrency.S1()
	concurrency.S2()
	concurrency.S3()
}
func TestForSelect(t *testing.T) {
	//	concurencypattern.Eg1()
	//	concurencypattern.Eg2()
	//	concurencypattern.Eg3()
	//	concurencypattern.Solve()
	concurencypattern.Solve2()
}
