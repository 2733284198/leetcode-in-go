package problem0059

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}

	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	num := 1

	mid := n>>1 + 1

	for i := 0; i < mid; i++ {
		tmpLen := n - 2*i - 1

		if tmpLen == 0 {
			result[i][i] = num
			continue
		}

		if tmpLen < 0 {
			break
		}

		for j := 0; j < tmpLen; j++ {
			result[i][i+j] = num
			num++
		}

		for j := 0; j < tmpLen; j++ {
			result[i+j][n-i-1] = num
			num++
		}

		for j := tmpLen; j > 0; j-- {
			result[n-i-1][i+j] = num
			num++
		}

		for j := tmpLen; j > 0; j-- {
			result[i+j][i] = num
			num++
		}

	}

	return result
}
