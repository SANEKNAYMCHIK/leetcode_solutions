func maxSubArray(nums []int) int {
	curSum := 0
	maxSum := math.MinInt
	for i := range nums {
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}