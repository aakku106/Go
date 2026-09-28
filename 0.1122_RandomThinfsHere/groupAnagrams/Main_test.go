package groupanagrams

import (
	"reflect"
	"testing"
)

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
		t.Run(v.name, func(t *testing.T) {
			if got := groupAnagrams(v.strs); !reflect.DeepEqual(got, v.expected) {
				t.Error("Expected: ", v.expected, "got: ", got)
			}
		})
	}
}
