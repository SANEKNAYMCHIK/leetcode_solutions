func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for i := range nums {
		if val, ok := data[target-nums[i]]; ok {
			return []int{val, i}
		}
		data[nums[i]] = i
	}
	return []int{}
}