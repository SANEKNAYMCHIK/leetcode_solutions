func rob(nums []int) int {
	n := len(nums)
	res := make([]int, n)
	res[0] = nums[0]
	if n > 1 {
		res[1] = max(nums[0], nums[1])
	}
	for i := 2; i < n; i++ {
		res[i] = max(res[i-2]+nums[i], res[i-1])
	}
	return res[n-1]
}