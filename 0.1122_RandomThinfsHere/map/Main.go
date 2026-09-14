package mapTimePass

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	S := make(map[byte]int8, len(s))
	T := make(map[byte]int8, len(S))
	// WE not usign ruine, insted using bytes type at key

	// Filling up the map
	for i, _ := range s {
		S[s[i]] = S[s[i]] + 1
		T[t[i]] = T[t[i]] + 1
	}
	for k, v := range S {
		if 
	}

	return false
}
