package examples

import (
	"testing"
)

func TestFiles(t *testing.T) {
	err := ReadFile("files.go")
	if err == nil {
		t.Fatal("The file ./files.go should exitst, but we get error")
	}
}
