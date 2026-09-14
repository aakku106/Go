package mapTimePass

func isAnagram(s string, t string) bool {
	// if both isen same than exit
	if len(s) != len(t) {
		return false
	}
	// Declearig 2 maps
	S := make(map[byte]int8, len(s))
	T := make(map[byte]int8, len(s))
	// WE not usign ruine, insted using bytes type at key

	// Filling up the map
	for i := range s {
		S[s[i]] = S[s[i]] + 1
		T[t[i]] = T[t[i]] + 1
	}

	// fmt.Println(S, "--S")
	// fmt.Println(T, "--T")

	count := 0
	for k := range S {
		if S[k] == T[k] {
			count++
		}
		// Checking if values of key in maps are same or not, and if yes how many of them, if all are same tham good else its not Anagram
	}
	if count == len(S) {
		return true
	}
	return false
}
