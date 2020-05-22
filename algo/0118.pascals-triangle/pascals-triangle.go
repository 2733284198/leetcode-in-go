package problem0118

func generate(numRows int) [][]int {
	var result [][]int

	if numRows == 0 {
		return result
	}

	result = append(result, []int{1})

	for i := 1; i < numRows; i++ {
		tmp := []int{0}
		tmp = append(tmp, result[i-1]...)

		for j := 0; j < len(tmp)-1; j++ {
			tmp[j] = tmp[j] + tmp[j+1]
		}

		result = append(result, tmp)
	}

	return result
}

func generateByDp(numRows int) [][]int {
	dp := make([][]int, numRows)
	if numRows == 0 {
		return dp
	}

	dp[0] = []int{1}

	for i := 1; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1 // 第一个值
		dp[i][i] = 1 // 最后一个值

		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}
