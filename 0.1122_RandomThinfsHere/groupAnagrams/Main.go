package groupanagrams

func groupAnagrams(strs []string) [][]string {
	first := strs[0]
}

// Function to take Anagrams, takes two strings and return true if anagram false else wise
func isAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//create a hash table
	Table := make(map[byte]int8, len(s))
	// fill the hash table
	for i := range s {
		Table[s[i]]++
		Table[t[i]]--
	}
	for key := range Table {
		if Table[key] != 0 {
			return false
		}
	}
	return true
}
