/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(node *TreeNode, data map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	leftPart := maxDepth(node.Left, data)
	rightPart := maxDepth(node.Right, data)
	res := max(leftPart, rightPart) + 1
	data[node] = res
	return res
}

func checkBBT(root *TreeNode, data map[*TreeNode]int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		tempRes := data[root.Left] - data[root.Right]
		if tempRes >= -1 && tempRes <= 1 {
			return checkBBT(root.Left, data) && checkBBT(root.Right, data)
		}
		return false
	}
	if root.Left == nil {
		if data[root.Right] <= 1 {
			return true
		}
		return false
	}
	if data[root.Left] <= 1 {
		return true
	}
	return false
}

func isBalanced(root *TreeNode) bool {
	data := make(map[*TreeNode]int)
	maxDepth(root, data)
	return checkBBT(root, data)
}