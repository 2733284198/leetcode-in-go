package problem0022

var num int

func generateParenthesis(n int) []string {

	result := make([]string, 0)

	backTracking(n, n, "", &result)

	return result
}

func backTracking(left, right int, tmp string, result *[]string) {

	// num++

	// fmt.Printf("num => %v left => %v right => %v tmp => %v result => %v\n", num, left, right, tmp, result)

	// 退出条件
	// 并不需要左括号是否用完，
	if right == 0 {
		*result = append(*result, tmp)
		// fmt.Println()
		return
	}

	// 生成左括号
	if left > 0 {
		backTracking(left-1, right, tmp+"(", result)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backTracking(left, right-1, tmp+")", result)
	}
}
