package problem0003

func lengthOfLongestSubstring(s string) int {
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	// 先设置所有的字符都没有见过
	for i := range location {
		location[i] = -1
	}
	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			//fmt.Println(s[left : i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i

		//fmt.Printf("location => %v \n", location)
		//fmt.Printf("s[i] => %v \n", s[i])
	}

	return maxLen
}
