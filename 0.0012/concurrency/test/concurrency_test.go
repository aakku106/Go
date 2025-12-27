package test

import (
	"concurrency"
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

func TestChannels(t *testing.T) {
	concurrency.T1()
	//	concurrency.T2()
	//
	// why i commented this t2 function is becaue will will block the main gorutine and waits till inifinit until someones kill the process
	// try Uncommenting if you dont get what am i trying to explain
	concurrency.T3()
}
