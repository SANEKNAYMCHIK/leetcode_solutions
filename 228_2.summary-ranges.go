func summaryRanges(nums []int) []string {
	var ans []string
	prevIdx := 0
	n := len(nums)
	if n == 0 {
		return ans
	}
	for i := 1; i < n; i++ {
		if (i - prevIdx) != (nums[i] - nums[prevIdx]) {
			if i-prevIdx == 1 {
				ans = append(ans, strconv.Itoa(nums[prevIdx]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[prevIdx], nums[i-1]))
			}
			prevIdx = i
		}
	}
	if n-prevIdx == 1 {
		ans = append(ans, strconv.Itoa(nums[prevIdx]))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[prevIdx], nums[n-1]))
	}
	return ans
}