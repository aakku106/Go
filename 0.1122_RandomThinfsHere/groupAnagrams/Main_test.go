package groupanagrams

import "testing"

func TestGroupAnagram(t *testing.T) {
	// strs := []string{"cat", "tac", "rat", "tar", "ate"}
	// groupAnagrams(strs)

	gg := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{"2 anagrams",
			[]string{"cat", "tac", "rat", "tar", "ate"},
			[][]string{{"cat", "tac"}, {"rat", "tar"}, {"ate"}},
		},
		{"only one",
			[]string{"cat"},
			[][]string{{"cat"}},
		},
	}

	for _, v := range gg {

	}
}
