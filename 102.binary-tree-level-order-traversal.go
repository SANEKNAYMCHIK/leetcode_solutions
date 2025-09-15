/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	vals := []*TreeNode{root}
	for len(vals) != 0 {
		vals2 := []*TreeNode{}
		tempVals := []int{}
		for len(vals) != 0 {
			curVal := vals[0]
			vals = vals[1:]
			if curVal == nil {
				continue
			}
			vals2 = append(vals2, curVal.Left)
			vals2 = append(vals2, curVal.Right)
			tempVals = append(tempVals, curVal.Val)
		}
		if len(tempVals) != 0 {
			ans = append(ans, tempVals)
		}
		for i := range vals2 {
			vals = append(vals, vals2[i])
		}
	}
	return ans
}