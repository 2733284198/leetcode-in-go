package problem0022

func generateParenthesis(n int) []string {

	result := make([]string, 0)

	backTracking(n, n, "", &result)

	return result
}

func backTracking(left, right int, tmp string, result *[]string) {

	if right == 0 {
		*result = append(*result, tmp)
		return
	}

	if left > 0 {
		backTracking(left-1, right, tmp+"(", result)
	}

	if right > left {
		backTracking(left, right-1, tmp+")", result)
	}
}
