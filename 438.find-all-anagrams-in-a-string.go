func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	ans := []int{}

	freqArr1 := make([]int, 26)
	freqArr2 := make([]int, 26)
	for i := range len(p) {
		freqArr1[p[i]-'a']++
		freqArr2[s[i]-'a']++
	}

	for i := 0; i < len(s)-len(p); i++ {
		if slices.Equal(freqArr1, freqArr2) {
			ans = append(ans, i)
		}
		freqArr2[s[i]-'a']--
		freqArr2[s[i+len(p)]-'a']++
	}
	if slices.Equal(freqArr1, freqArr2) {
		ans = append(ans, len(s)-len(p))
	}
	return ans
}