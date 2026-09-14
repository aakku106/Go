package mapTimePass

import (
	"fmt"
	"testing"
)

var (
	a = "cat"
	b = "tac"
	c = "cati"
	d = "taci"
	e = "camera"
	f = "ccmera"
)

func TestIsAnagram(t *testing.T) {

	if ab := isAnagram(a, b); ab == false {
		t.Error("cat amd tac is Anagram")
	} else {
		fmt.Println(a, b, "Was Good")
	}

	if bc := isAnagram(b, c); bc == true {
		t.Error("cati and tac is not Anagram")
	} else {
		fmt.Println(b, c, "Was Not Good")
	}

	if cd := isAnagram(d, c); cd == false {
		t.Error(" cati and taci is Anagram")
	} else {
		fmt.Println(c, d, "Was Good")
	}

	if de := isAnagram(d, e); de == true {
		t.Error(" taci and camera isn't Anagram")
	} else {
		fmt.Println(d, e, "Wasn't Good")
	}

	if ef := isAnagram(e, f); ef == true {
		t.Error(" ccmera and camera isn't Anagram")
	} else {
		fmt.Println(e, f, "Wasn't Good")
	}
}
