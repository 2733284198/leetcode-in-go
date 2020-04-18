package problem0151

func reverseWords(str string) string {
	s := []byte(str)
	// 去掉多余的空格
	// for i := 1; i < len(s); i++ {
	// 	if s[i] == ' ' && s[i-1] == ' ' {
	// 		s = append(s[0:i-1], s[i:])
	// 	}
	// }

	length := len(s)

	// 先反转整个字符串
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	// 在将单词反转回来
	start, end := 0, 0
	for i := 0; i < length; i++ {
		if s[i] == ' ' || i == length-1 { // 最后一个单词反转
			if i == length-1 {
				end = i
			} else {
				end = i - 1
			}
			reverseString(s, start, end)
			start = i + 1 // 得到下一个单词的开始
		}
	}

	return string(s)
}

func reverseString(s []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWordsV2(s string) string {
	if s == "" {
		return ""
	}
	res := []byte{}
	queue := []string{}

	word := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				queue = append(queue, string(word))
				word = []byte{}
			}
		} else {
			word = append(word, s[i])
		}
	}
	if len(word) > 0 {
		queue = append(queue, string(word))
	}

	if len(queue) <= 0 {
		return ""
	}

	for i := len(queue) - 1; i >= 0; i-- {
		res = append(res, []byte(queue[i])...)
		res = append(res, ' ')
	}

	return string(res[:len(res)-1])
}
