func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	vals := make(map[string][]string)
	for i := range strs {
		newVal := strings.Split(strs[i], "")
		sort.Strings(newVal)
		key := strings.Join(newVal, "")
		vals[key] = append(vals[key], strs[i])
	}
	for _, value := range vals {
		ans = append(ans, value)
	}
	return ans
}
