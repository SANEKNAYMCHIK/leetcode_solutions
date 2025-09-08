func longestSubarray(nums []int) int {
	stringNums := fmt.Sprint(nums)
	countsZero := strings.Contains(stringNums, "0")
	stringVals := strings.Split(strings.Trim(stringNums, "[]"), " ")
	newStringVals := strings.Split(strings.Join(stringVals, ""), "0")
	if countsZero {
		ans := 0
		for i := 1; i < len(newStringVals); i++ {
			ans = max(ans, len(newStringVals[i-1])+len(newStringVals[i]))
		}
		return ans
	} else {
		return len(newStringVals[0]) - 1
	}
}
