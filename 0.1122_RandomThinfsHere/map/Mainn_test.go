package mapTimePass

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	a := "cat"
	b := "tac"
	c := "cati"
	d := "taci"

	if ab := isAnagram(a, b); ab == false {
		t.Error("cat amd tac is Anagram")
	} else {
		fmt.Println(a, b, "Was Good")
	}

	if bc := isAnagram(b, c); bc == true {
		t.Error("cati and tac is not Anagram")
	} else {
		fmt.Println(a, b, "Was Good")
	}

	if cd := isAnagram(c, d); cd == false {
		t.Error(" cati and taci is Anagram")
	} else {
		fmt.Println(a, b, "Was Good")
	}

}
