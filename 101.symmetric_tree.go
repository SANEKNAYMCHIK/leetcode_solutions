/*
	Definition for a binary tree node.
	type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}
*/

func sym(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return sym(left.Left, right.Right) && sym(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return sym(root.Left, root.Right)
}