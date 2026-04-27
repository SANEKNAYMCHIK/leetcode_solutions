func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	data := make(map[string]bool)
	dp := make([]bool, n+1)
	dp[0] = true
	maxLen := 0
	var curLen int
	for _, val := range wordDict {
		curLen = len(val)
		if curLen > maxLen {
			maxLen = curLen
		}
		data[val] = true
	}
	for i := 1; i < n+1; i++ {
		for j := i - 1; j >= max(i-maxLen-1, 0); j-- {
			if dp[j] && data[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}