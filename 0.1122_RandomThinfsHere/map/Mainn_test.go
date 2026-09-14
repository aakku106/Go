package mapTimePass

import "testing"

func TestisAnagram(t *testing.T) {
	a := "cat"
	b := "tac"
	c := "cati"
	d := "taci"

	if ab := isAnagram(a, b); ab == false {
		t.Error("It cat adn tac is Anagram")
	}

	if bc := isAnagram(b, c); bc == true {
		t.Error("It cati adn tac is not Anagram")
	}
	if cd := isAnagram(c, d); cd == false {
		t.Error("It cati and taci is Anagram")
	}

}
