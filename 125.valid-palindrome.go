func isPalindrome(s string) bool {
	var res []rune
	for _, i := range s {
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			res = append(res, unicode.ToLower(i))
		}
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		if res[i] != res[n-i-1] {
			return false
		}
	}
	return true
}
