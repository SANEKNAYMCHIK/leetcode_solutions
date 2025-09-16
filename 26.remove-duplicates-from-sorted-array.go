func removeDuplicates(nums []int) int {
	ans := 1
	n := len(nums)
	i, j := 0, 0
	for j < n {
		if nums[j] != nums[i] {
			ans++
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return ans
}