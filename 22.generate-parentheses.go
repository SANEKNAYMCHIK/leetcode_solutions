func dfs(op, cl int, val string, ans *[]string) {
	if (op == 0) && (cl == 0) {
		(*ans) = append((*ans), val)
	}
	if op < cl {
		dfs(op, cl-1, val+")", ans)
	}
	if op > 0 {
		dfs(op-1, cl, val+"(", ans)
	}
}

func generateParenthesis(n int) []string {
	open := n
	close := n
	ans := []string{}
	dfs(open, close, "", &ans)
	return ans
}
