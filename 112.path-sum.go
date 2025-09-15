/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findPath(node *TreeNode, curVal, needVal int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		if curVal+node.Val == needVal {
			return true
		}
		return false
	}
	return findPath(node.Left, curVal+node.Val, needVal) || findPath(node.Right, curVal+node.Val, needVal)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return findPath(root, 0, targetSum)
}