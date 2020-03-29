package problem0123

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	maxK := 2 // max trade nums

	dp := make([][][]int, size)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j <= maxK; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	// 第 0 天 持有股票
	dp[0][0][1] = -prices[0]
	// 第 0 天持有股票，交易 1 比
	dp[0][1][1] = -prices[0]
	// 第 0 天持有股票，交易 2 比
	dp[0][2][1] = -prices[0]

	for i := 1; i < size; i++ {
		for k := 1; k <= maxK; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[size-1][maxK][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitV2(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	profits := []int{}
	temp := 0
	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]

		if temp*diff >= 0 {
			temp += diff
			continue
		}

		profits = append(profits, temp)
		temp = diff
	}
	profits = append(profits, temp)

	res := 0
	for i := 0; i < len(profits); i++ {
		temp = maxSliceEle(profits[:i]) + maxSliceEle(profits[i:])
		if res < temp {
			res = temp
		}
	}

	return res
}

func maxSliceEle(ps []int) int {
	max, tmp := 0, 0

	for _, p := range ps {
		if tmp < 0 {
			tmp = 0
		}

		tmp += p

		if max < tmp {
			max = tmp
		}
	}

	return max
}

// func maxProfitV3(prices []int) int {
// 	if len(prices) <= 1 {
// 		return 0
// 	}

// 	s1, s2, s3, s4 := prices[0], math.MinInt32, math.MinInt32, math.MinInt32

// 	for i := 1; i < len(prices); i++ {
// 		s1 = max(s1, -prices[i])
// 		s2 = max(s2, s1+prices[i])
// 		s3 = max(s3, s2-prices[i])
// 		s4 = max(s4, s3+prices[i])
// 	}

// 	return max(0, s4)
// }
