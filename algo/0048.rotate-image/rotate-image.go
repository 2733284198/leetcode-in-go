package problem0048

func rotate(matrix [][]int) {
	// i,j 翻转90度变为 j,n-1-i
	n := len(matrix)
	for i := 0; i <= n/2; i++ {
		for j := i; j <= n-i-2; j++ {
			// matrix[i][j] =  matrix[n-j-1][i]
			// matrix[j][n-i-1] = matrix[i][j]
			// matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			// matrix[n-j-1][i] = matrix[n-i-1][n-j-1]

			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}

	return
}

func rotateV2(matrix [][]int) {
	mLen := len(matrix)

	if mLen == 0 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}

	var temp int // tmp value store

	for i := 0; i < mLen/2; i++ {
		for j := i; j < mLen-i-1; j++ {
			temp = matrix[mLen-1-i][j]

			matrix[mLen-1-i][j] = matrix[mLen-1-j][mLen-1-i]
			matrix[mLen-1-j][mLen-1-i] = matrix[i][mLen-1-j]
			matrix[i][mLen-1-j] = matrix[j][i]

			matrix[j][i] = temp
		}
	}
}
