package problem0239

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || nums == nil {
		return []int{}
	}

	// window 存储 nums 的下标
	// result 存储每次遍历中 window 得到最大值
	window, result := []int{}, []int{}

	for i, v := range nums {

		// 维护 window 元素的数量
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		// 删除 window 列表最后一个元素，如果小于V
		for len(window) > 0 && nums[window[len(window)-1]] <= v {
			window = window[:len(window)-1]
		}

		// 窗口向右移动一次
		window = append(window, i)

		if i >= k-1 { // i 的下标是从 0 开始的
			result = append(result, nums[window[0]])
		}
	}

	return result
}
