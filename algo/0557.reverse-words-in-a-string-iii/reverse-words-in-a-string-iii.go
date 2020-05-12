package problem0557

func reverseWords(s string) string {
	b := []byte(s)
	l, n := 0, len(s)

	for i, v := range s {
		if v == ' ' ||  i == n -1 {
			r := i - 1
			if i == n - 1 { // 最后一个单次没有空格
				r = i
			}

			for l < r {
				b[l], b[r] = b[r], b[l]
				l ++
				r --
			}
			l = i + 1
		} 	
	}

	return string(b)
}