package problem0036

func isValidSudoku(board [][]byte) bool {
	coloum := make([]int, 9*9) //9列，每行9个数
	row := make([]int, 9*9)    //9行，每行9个数
	box := make([]int, 9*9)    //9小格，每行9个数
	result := true

	boxIdx := func(i, j, n int) int {
		return (i/3+(j/3)*3)*9 + n
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := int(board[i][j] - '1')
			if coloum[j*9+n] == 1 || row[i*9+n] == 1 || box[boxIdx(i, j, n)] == 1 {
				return false
			}
			coloum[j*9+n] = 1
			row[i*9+n] = 1
			box[boxIdx(i, j, n)] = 1
		}
	}

	return result
}
