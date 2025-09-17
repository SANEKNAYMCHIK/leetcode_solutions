func subarraysDivByK(nums []int, k int) int {
	ans := 0
	data := make(map[int]int)
	data[0] = 1
	prefSum := 0
	for i := range nums {
		prefSum = (prefSum + nums[i]) % k
		if prefSum < 0 {
			prefSum += k
		}
		if val, ok := data[prefSum]; ok {
			ans += val
		}
		data[prefSum]++
	}
	return ans
}