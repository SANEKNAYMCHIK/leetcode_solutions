func lengthOfLIS(nums []int) int {
	res := []int{nums[0]}
	n := 0
	lenNums := len(nums)

	binSearch := func(val int) int {
		low, high := 0, n+1
		for low < high {
			med := low + (high-low)/2
			if res[med] < val {
				low = med + 1
			} else {
				high = med
			}
		}
		return low
	}

	for i := 1; i < lenNums; i++ {
		if nums[i] > res[n] {
			res = append(res, nums[i])
			n++
		} else {
			res[binSearch(nums[i])] = nums[i]
		}
	}
	return len(res)
}