func summaryRanges(nums []int) []string {
	var ans []string
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		ans = append(ans, fmt.Sprintf("%d", nums[0]))
		return ans
	}
	startItem := nums[0]
	prevItem := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-prevItem == 1 {
			prevItem = nums[i]
		} else {
			if startItem == prevItem {
				ans = append(ans, fmt.Sprintf("%d", startItem))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", startItem, prevItem))
			}
			startItem = nums[i]
			prevItem = nums[i]
		}
	}
	if startItem == prevItem {
		ans = append(ans, fmt.Sprintf("%d", startItem))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", startItem, prevItem))
	}
	return ans
}
