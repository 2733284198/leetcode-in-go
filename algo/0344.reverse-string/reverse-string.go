package problem0344

func reverseString(str string) string {
	s := []byte(str)
	// l := len(s)
	// mid := l / 2
	// for i := 0; i < mid; i++ {
	// 	s[i], s[l-i-1] = s[l-i-1], s[i]
	// }

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
