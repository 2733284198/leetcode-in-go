package problem0052

var res [][]string

func totalNQueens(n int) int {
	res = nil
	pos := make([]int, n+1)
	dfs(pos, 1, n)

	return len(res)
}

func dfs(pos []int, row, n int) {
	if row > n {
		res = append(res, genBoard(pos, n))
		return
	}

	for col := 1; col <= n; col++ {
		valid := true

		for i := 1; i <= n; i++ {
			if row == i || pos[i] == 0 {
				continue
			}

			if col == pos[i] { // on the same col
				valid = false
				break
			}

			if abs(row-i) == abs(col-pos[i]) { // on the same diagonal
				valid = false
				break
			}
		}

		if valid {
			pos[row] = col
			dfs(pos, row+1, n)
			pos[row] = 0 // reset this row
		}
	}

}

func genBoard(pos []int, n int) (result []string) {
	for i := 1; i <= n; i++ {
		col := pos[i]
		rowStr := ""
		for j := 1; j <= n; j++ {
			if j == col {
				rowStr += "Q"
			} else {
				rowStr += "."
			}
		}
		result = append(result, rowStr)
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
