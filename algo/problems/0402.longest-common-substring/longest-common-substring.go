package problems0402

func maxSubTwoString(str1, str2 string) string {
	chs1 := len(str1)
	chs2 := len(str2)

	maxLength := 0 //记录最大长度
	end := 0       //记录最大长度的结尾位置

	rows := 0
	cols := chs2 - 1
	for rows < chs1 {
		i, j := rows, cols
		length := 0 //记录长度
		for i < chs1 && j < chs2 {
			if str1[i] != str2[j] {
				length = 0
			} else {
				length++
			}
			if length > maxLength {
				end = i
				maxLength = length
			}
			i++
			j++
		}
		if cols > 0 {
			cols--
		} else {
			rows++
		}
	}

	//fmt.Printf("str1 => %v", str1)

	// return false

	return str1[(end - maxLength + 1):(end + 1)]
}

func longestCommonSubString(text1, text2 string) string {
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	end, maxLen := 0, 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				end = i
				maxLen = max(maxLen, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}

	//fmt.Println(dp, end, maxLen)

	return text1[end-maxLen : end]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
