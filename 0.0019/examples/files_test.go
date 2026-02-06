package examples

import (
	"testing"
)

func TestFiles(t *testing.T) {
	err := ReadFile("files.go")
	if err != nil {
		t.Fatal("The file ./files.go should exitst, but we get error")
	}

	err = ReadFile2("files.go")
	if err != nil {
		t.Fatal("The file ./files.go should exitst, but we get error")
	}

}
func TestSomeWork(t *testing.T) {

	someWork()
	someBetterWork()
	SomeFileReading("./go.mod")
	SomeFileReading("./go.sum")

}
