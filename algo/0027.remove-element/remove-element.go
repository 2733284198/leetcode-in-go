package problem0027

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			// nums[i] 已经被删除 保持下标不变 抵消掉循环中的 i++
		}
	}

	return len(nums)
}
