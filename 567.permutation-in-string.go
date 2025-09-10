func check(freq1, freq2 []int) bool {
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	freqArr1 := make([]int, 26)
	freqArr2 := make([]int, 26)
	for i := range s1 {
		freqArr1[s1[i]-'a']++
		freqArr2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if check(freqArr1, freqArr2) {
			return true
		}
		freqArr2[s2[i]-'a']--
		freqArr2[s2[i+len(s1)]-'a']++
	}
	return check(freqArr1, freqArr2)
}