package problem0073

func setZeroes(matrix [][]int) {
	//判断第1列是否有0
	row0_zero, col0_zero := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0_zero = true
			break
		}
	}
	//判断第1行是否有0
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			row0_zero = true
			break
		}
	}

	//判断从第2行和第2列开始，若出现0，则在第一行和第一列上分别做标记
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	//遍历除第一行和第一列外的所有元素，若对应第1行或第1列的位置为0，则该matrix[i][j]位置也置0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	//若之前判断第1列有0，则将整个第1列置零
	if col0_zero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	//若之前判断第1行有0，则将整个第1行置零
	if row0_zero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}
