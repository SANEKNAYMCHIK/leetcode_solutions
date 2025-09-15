/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	var getPath func(node *TreeNode, curVal, needVal int, steps []int)
	getPath = func(node *TreeNode, curVal, needVal int, steps []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if curVal+node.Val == needVal {
				steps = append(steps, node.Val)
				newVal := []int{}
				for i := range steps {
					newVal = append(newVal, steps[i])
				}
				ans = append(ans, newVal)
			}
			return
		}
		steps = append(steps, node.Val)
		getPath(node.Left, curVal+node.Val, needVal, steps)
		getPath(node.Right, curVal+node.Val, needVal, steps)
	}
	getPath(root, 0, targetSum, []int{})
	return ans
}