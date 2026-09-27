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
	var final [][]string

mainLoop:
	for _, value := range strs {
		if len(value) > 100 {
			return nil
		}
		if value == first {
			one = append(one, first)
			continue mainLoop
		}
		ok := isAnagrams(first, value)
		if ok {
			one = append(one, value)
			// fmt.Println(value, "Enterde in One")
		} else {
			notAnargamList = append(notAnargamList, value)
			// fmt.Println(value, " Entered in NotAnargam List")
		}

	}
	final = append(final, one)
	// fmt.Println(one, " Entered in Final")
	one = []string{}
	// fmt.Println("After MainLoop not= ", notAnargamList)
	// fmt.Println("Main Loop done, Final= ", final)
	newNotAnargamsList := []string{}

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
		newNotAnargamsList = []string{}
		one = []string{}
		// fmt.Println(one, " Entered in Final--")
	}
	// fmt.Println(final, "That's Final")
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
