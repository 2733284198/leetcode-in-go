package problem0055

func canJump(nums []int) bool {
	max := 0

	// 遍历数组，得到当前位置最长能跳多远max，碰到0看看max能跳过去不，不能跳过去就是false，注意边界条件。
	for index, n := range nums {
		if n == 0 && index != len(nums)-1 {
			if max <= 1 {
				return false
			}
		}
		max = max - 1
		if n > max {
			max = n
		}
	}

	return true
}
