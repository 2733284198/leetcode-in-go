package problem0119

func getRow(rowIndex int) []int {
	nums := []int{1} // 第0行
	for i := 1; i <= rowIndex; i++ {
		nums = append(nums, 1) // 尾部追加1
		for j:=i-1;j>0;j--{
			nums[j]= nums[j] + nums[j-1]
		}
	}
	return nums
}

func getRowByMap(rowIndex int) []int {
	result := generateByDp(rowIndex + 1)

	return result[rowIndex]
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
