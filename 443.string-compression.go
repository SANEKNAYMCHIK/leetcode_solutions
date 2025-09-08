func compress(chars []byte) int {
	counter := 1
	val := chars[0]
	for i := 1; i < len(chars); i++ {
		if val != chars[i] {
			if counter != 1 {
				curVal := strconv.Itoa(counter)
				k := 0
				n := len(curVal)
				for j := i - counter + 1; j < i; j++ {
					if k < n {
						chars[j] = curVal[k]
						k++
					} else {
						chars[j] = 0
					}
				}
			}
			counter = 1
			val = chars[i]
		} else if (val == chars[i]) && (i == len(chars)-1) {
			counter++
			curVal := strconv.Itoa(counter)
			k := 0
			n := len(curVal)
			for j := i - counter + 2; j < i+1; j++ {
				if k < n {
					chars[j] = curVal[k]
					k++
				} else {
					chars[j] = 0
				}
			}
			counter = 1
			val = chars[i]
		} else {
			counter++
		}
	}
	i := 0
	for {
		if (chars[i] == 0) && (i < len(chars)-1) {
			chars = append(chars[:i], chars[i+1:]...)
		} else if (chars[i] == 0) && (i == len(chars)-1) {
			chars = chars[:i]
		} else {
			i++
		}
		if i >= len(chars) {
			break
		}
	}
	return len(chars)
}