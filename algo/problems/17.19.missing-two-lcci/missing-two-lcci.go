package problems1719

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	S0 := ((1 + n) * n) / 2
	S1, a, b, mid := 0, 0, 0, 0
	for _, num := range nums {
		S1 += num
	}
	mid = (S0 - S1) / 2
	midTotal := ((1 + mid) * mid) / 2
	for _, num := range nums {
		if num <= mid {
			midTotal = midTotal - num
		}
	}

	a = midTotal
	b = S0 - S1 - a

	return []int{a, b}
}

func missingTwoV2(nums []int) []int {
	nums = append(nums, 1, 2)          // 补两个数
	for i := 0; i < len(nums)-2; i++ { // 将 原来的 nums 元素设置为复数
		idx := nums[i]
		if idx < 0 {
			idx = -idx
		}
		idx -= 1
		nums[idx] = -nums[idx]
	}

	var res []int
	for i, num := range nums {
		if num >= 0 {
			res = append(res, i+1)
		}
	}

	return res
}
