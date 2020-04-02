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
