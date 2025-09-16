func subarraySum(nums []int, k int) int {
	data := make(map[int]int)
	data[0] = 1
	sumPref := 0
	ans := 0
	for i := range nums {
		sumPref += nums[i]
		if val, ok := data[sumPref-k]; ok {
			ans += val
		}
		data[sumPref]++
	}
	return ans
}