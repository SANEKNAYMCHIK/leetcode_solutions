/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var getLevel func(node *TreeNode, level int)
	getLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ans) == level {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], node.Val)
		getLevel(node.Left, level+1)
		getLevel(node.Right, level+1)
	}
	getLevel(root, 0)
	for i := range ans {
		if i%2 != 0 {
			slices.Reverse(ans[i])
		}
	}
	return ans
}