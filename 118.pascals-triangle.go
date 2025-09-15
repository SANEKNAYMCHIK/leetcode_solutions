func generate(numRows int) [][]int {
	ans := [][]int{}
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	for i := 2; i < numRows; i++ {
		newVals := []int{1}
		for j := 0; j < len(ans[i-1])-1; j++ {
			newVals = append(newVals, (ans[i-1][j] + ans[i-1][j+1]))
		}
		newVals = append(newVals, 1)
		ans = append(ans, newVals)
	}
	return ans
}