func lengthOfLongestSubstring(s string) int {
	letters := make(map[rune]int)
	ans := 0
	for i, j := 0, 0; j < len(s); {
		if _, ok := letters[rune(s[j])]; ok {
			i = max(letters[rune(s[j])]+1, i)
			letters[rune(s[j])] = j
		} else {
			letters[rune(s[j])] = j
		}
		ans = max(ans, j-i+1)
		j++
	}
	return ans
}
