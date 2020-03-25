package problem0069

func mySqrt(x int) int {
	// 使用 math.Sqrt 库函数
	// 使用牛顿法求平方根
	r := x

	for r*r > x {
		r = (r + x/r) / 2
	}

	return r
}
