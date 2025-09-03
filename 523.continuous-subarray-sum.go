func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	rems := make(map[int]int)
	prefSum := 0
	for i := range len(nums) {
		prefSum += nums[i]
		rem := prefSum % k
		if rem == 0 {
			if i+1 >= 2 {
				return true
			} else {
				rems[rem] = i
			}
		} else {
			if val, ok := rems[rem]; ok {
				if i-val >= 2 {
					return true
				}
			} else {
				rems[rem] = i
			}
		}
	}
	return false
}
