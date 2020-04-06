package problem0032

func longestValidParentheses(s string) int {
	var left, max, temp int
	record := make([]int, len(s))

	//1. Record
	for i, b := range s {
		if b == '(' {
			left++
		} else if left > 0 {
			left--
			record[i] = 2
		}
	}

	//2. update record
	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
	}

	//3. result
	for _, r := range record {
		if r == 0 {
			temp = 0
			continue
		}

		temp++
		if temp > max {
			max = temp
		}
	}

	return max
}

func longestValidParenthesesByStack(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	result := 0
	dp := make([]int, n) // dp[i] 表示已经 s[i] 结尾的最长有效括号
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == ')' && len(stack) != 0 {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			dp[i] = i - idx + 1
			if idx-1 >= 0 {
				dp[i] = dp[i] + dp[idx-1]
			}
			result = max(result, dp[i])
		}

		if s[i] == '(' {
			stack = append(stack, i)
		}
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
