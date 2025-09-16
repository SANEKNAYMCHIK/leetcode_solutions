func topKFrequent(words []string, k int) []string {
	res := make(map[string]int)
	vals := []string{}
	for i := range words {
		if _, ok := res[words[i]]; !ok {
			vals = append(vals, words[i])
		}
		res[words[i]]++
	}
	sort.Slice(vals, func(i, j int) bool {
		if res[vals[i]] == res[vals[j]] {
			return vals[i] < vals[j]
		}
		return res[vals[j]] < res[vals[i]]
	})
	// slices.SortFunc(vals, func(i, j string) int {
	// 	if res[i] == res[j] {
	// 		return strings.Compare(strings.ToLower(i), strings.ToLower(j))
	// 	}
	// 	return res[j] - res[i]
	// })
	return vals[:k]
}
