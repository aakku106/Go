package test

import (
	"concurrency"
	"fmt"
	"testing"
)

func TestConc(t *testing.T) {
	fmt.Println("weeee")
	concurrency.C1()
}
