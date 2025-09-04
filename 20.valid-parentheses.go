func isValid(s string) bool {
	vals := []rune{}
	for _, item := range s {
		if (item == '(') || (item == '[') || (item == '{') {
			vals = append(vals, item)
		} else {
			n := len(vals) - 1
			if (n > -1) && ((item == ')' && vals[n] == '(') || (item == ']' && vals[n] == '[') || (item == '}' && vals[n] == '{')) {
				vals = vals[:n]
			} else {
				return false
			}
		}
	}
	if len(vals) != 0 {
		return false
	}
	return true
}