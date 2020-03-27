package problem0070

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	var first, second, third = 1, 2, 0

	for i := 3; i <= n; i++ {
		third = first + second
		first = second
		second = third
	}

	return second
}

func climbStairsByBF(n int) int {
	return climbBruteForce(0, n)
}

func climbBruteForce(i, n int) int  { // i 定义了当前阶数，n 定义了目标阶数
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return climbBruteForce(i + 1, n) + climbBruteForce(i + 2, n)
}

func climbStairsDp(n int) int{
	if n == 1 {
		return 1
	}

	dp := make([]int, n + 1)

	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}