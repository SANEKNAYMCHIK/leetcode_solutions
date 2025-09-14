func countDepth(node *TreeNode, counter int) int {
	if node == nil {
		return counter
	}
	return max(countDepth(node.Right, counter+1), countDepth(node.Left, counter+1))
}

func maxDepth(root *TreeNode) int {
	return countDepth(root, 0)
}