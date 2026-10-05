package groupanagrams

// import "fmt"

func groupAnagrams(strs []string) [][]string {
	// fmt.Println("Given list: ", strs)
	if len(strs) == 0 || len(strs) > 10000 {
		return nil
	}

	first := strs[0]
	one := make([]string, 0, len(strs)/2)
	notAnargamList := make([]string, 0, len(strs)/2)

	for _, value := range strs {
		if len(value) > 100 {
			return nil
		}
		if value == first {
			one = append(one, first)
			continue
		}
		if ok := isAnagrams(first, value); ok {
			one = append(one, value)
			// fmt.Println(value, "Enterde in One")
		} else {
			notAnargamList = append(notAnargamList, value)
			// fmt.Println(value, " Entered in NotAnargam List")
		}
	}

	var final [][]string
	final = append(final, one)
	// fmt.Println(one, " Entered in Final")
	one = make([]string, 0, len(notAnargamList)/2)
	// fmt.Println("After MainLoop not= ", notAnargamList)
	// fmt.Println("Main Loop done, Final= ", final)
	newNotAnargamsList := make([]string, 0, len(notAnargamList)/2)

	for len(notAnargamList) != 0 {
		if len(notAnargamList) == 1 {
			final = append(final, notAnargamList)
			// fmt.Println("Found single one", notAnargamList)
			break
		}
		first = notAnargamList[0]
		// fmt.Println("Entering notAnargams List")
		for _, value := range notAnargamList {
			if value == first {
				one = append(one, first)
				continue
			}

			if ok := isAnagrams(first, value); ok {
				one = append(one, value)
				// fmt.Println(value, "Enterde in One")
			} else {
				newNotAnargamsList = append(newNotAnargamsList, value)
				// fmt.Println(value, " Entered in NotAnargam List")
			}
		}
		final = append(final, one)
		notAnargamList = newNotAnargamsList
		newNotAnargamsList = make([]string, 0, len(notAnargamList)/2)
		one = make([]string, 0, len(notAnargamList)/2)
		// fmt.Println(one, " Entered in Final--")
	}
	// fmt.Println(final, "That's Final")
	return final
}

func groupAnagramsImproved(strs []string) [][]string {

	if len(strs) == 0 || len(strs) > 10000 {
		return [][]string{}
	}
	final := make([][]string, 0, len(strs)/2)
	notAnargamGroup := strs

	for len(notAnargamGroup) != 0 {
		if len(notAnargamGroup) == 1 {
			final = append(final, notAnargamGroup)
			break
		}
		first := notAnargamGroup[0]
		newNotAnargamsList := make([]string, 0, len(notAnargamGroup)/2)
		anagramGroup := make([]string, 0, len(strs)/2)
		for _, v := range notAnargamGroup {
			if v == first {
				anagramGroup = append(anagramGroup, v)
				continue
			}
			if ok := isAnagrams(first, v); ok {
				anagramGroup = append(anagramGroup, v)
			} else {
				newNotAnargamsList = append(newNotAnargamsList, v)
			}
		}
		final = append(final, anagramGroup)
		notAnargamGroup = newNotAnargamsList
	}

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

func groupAnagramsOptimized(strs []string) [][]string {
	if len(strs) == 0 || len(strs) > 10000 {
		return [][]string{}
	}

	g := make(map[[26]byte]string)

	return nil
}
