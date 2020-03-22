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

func totalNQueensV2(n int) int {
	res := 0
	// 生成一个n x n的数组
	board := make([][]rune, n)
	for i := 0; i < n; i++ {
		item := make([]rune, n)
		for k, _ := range item {
			item[k] = '.'
		}
		board[i] = item
	}

	backTrack(board, &res, 0)
	return res
}

func backTrack(board [][]rune, res *int, row int) {
	// 退出条件
	if row == len(board) {
		*res++
		return
	}

	// N叉树遍历
	for col := 0; col < len(board); col++ {
		// 排除不合法的数据
		if isOk(board, row, col) {
			// 做选择
			board[row][col] = 'Q'
			// 进入下一层
			backTrack(board, res, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
}

func isOk(board [][]rune, row, col int) bool {

	// 检查上方数据
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左上方数据
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 检查右上方
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
