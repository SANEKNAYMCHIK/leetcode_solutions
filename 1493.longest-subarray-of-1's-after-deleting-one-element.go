func longestSubarray(nums []int) int {
	var counterZeros int
	var counterOnes int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			counterZeros += 1
		} else {
			counterOnes += 1
		}
	}
	if counterOnes == 0 {
		return 0
	}
	if counterZeros == 0 {
		return len(nums) - 1
	}
	var tempSubArrs []int
	tempLen := 0
	maxVal := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tempLen += 1
			if i == len(nums)-1 {
				tempSubArrs = append(tempSubArrs, tempLen)
				maxVal = max(maxVal, tempLen)
			}
		} else {
			if tempLen != 0 {
				maxVal = max(maxVal, tempLen)
				tempSubArrs = append(tempSubArrs, tempLen)
				tempLen = 0
			} else {
				if i != 0 {
					tempSubArrs = append(tempSubArrs, 0)
				}
			}
		}
	}
	if len(tempSubArrs) == 1 {
		return maxVal
	}
	maxLen := tempSubArrs[0] + tempSubArrs[1]
	for i := 2; i < len(tempSubArrs); i++ {
		if tempSubArrs[i-1]+tempSubArrs[i] > maxLen {
			maxLen = tempSubArrs[i-1] + tempSubArrs[i]
		}
	}
	return maxLen
}
