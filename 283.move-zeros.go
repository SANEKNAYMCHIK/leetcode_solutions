func moveZeroes(nums []int) {
	l := 0
	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
}
