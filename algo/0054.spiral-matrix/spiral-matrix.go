package problem0054

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 || matrix == nil {
		return result
	}

	m, n := len(matrix), len(matrix[0]) // m 行 n 列
	up, down, left, right := 0, m-1, 0, n-1

	for {
		// left -> right
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		if up++; up > down {
			break
		}

		// up -> down
		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		if right--; right < left {
			break
		}

		// right -> left
		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}
		if down--; down < up {
			break
		}

		// down -> up
		for i := down; i >= up; i-- {
			result = append(result, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}

	return result
}
