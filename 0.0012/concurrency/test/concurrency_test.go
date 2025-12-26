package test

import (
	"concurrency"
	"testing"
)

func TestConc(t *testing.T) {
	// concurrency.WaitToGetGet()
	concurrency.C1()
	// we won't see any output from go get function of C1, function (but those get function will be printed because we waited for some seconds in WaitToGetGet functer below, thats the resion all go get from c1 wont be printed cause timwA we waited weren't enougth to print them all)
	concurrency.WaitToGetGet()
	// if you call  concurrency.WaitToGetGet() before concurrency.C1() you can observe only wee will be printed
}
