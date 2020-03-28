package problem0121

import "math"

func maxProfit(prices []int) int {
	min, max := math.MaxInt32, 0
	// i j (j > i) max(prices[j] - prices[i])

	for i := 0; i < len(prices); i++ {
		if prices[i] < min { // buy
			min = prices[i]
		}
		if math.MaxInt32 != min && prices[i]-min > max { // sell
			max = prices[i] - min
		}
	}

	return max
}
