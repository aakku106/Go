package groupanagrams

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	first := strs[0]
	one := make([]string, 0, len(strs)/2)
	notAnargamList := make([]string, 0, len(strs)/2)
	var final [][]string

mainLoop:
	for _, value := range strs {
		if value == first {
			break mainLoop
		}
		ok := isAnagrams(first, value)
		if ok {
			one = append(one, value)
		} else {
			notAnargamList = append(notAnargamList, value)
		}

	}
	final = append(final, one)
	one = []string{}

anotherLoop:
	for len(notAnargamList) != 0 {
		first = notAnargamList[0]
		for _, value := range notAnargamList {
			if value == first {
				break anotherLoop
			}
			if len(notAnargamList) == 1 {
				final = append(final, notAnargamList)
			}
			if ok := isAnagrams(first, value); ok {
				one = append(one, value)
			} else {
				notAnargamList = append(notAnargamList, value)
			}
		}
	}

	final = append(final, one)
	return final
}

// Function to check Anagrams, takes two strings and return true if anagram false else wise
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
	// check if 0, cause repeat=0
	for key := range Table {
		if Table[key] != 0 {
			return false
		}
	}
	return true
}
