package problem0209

func minSubArrayLen(s int, nums []int) int  {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 定义滑动窗口的开始、结束指针 i、j
	i,j,sum := 0,0, nums[0]
	minLen := len(nums) + 1
	// 窗口大小就是要求的数组长度 一直维持滑动窗口的和 >= s
	// 如果小了 右边加元素 j ++
	// 如果大了 左边减元素 i ++
	for j < len(nums) {
		if sum >= s {
			if j - i +1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i ++
		} else {
			j ++
			if j >= len(nums) {
				break
			}
			sum += nums[j]
		}
	}

	if minLen == len(nums) + 1 {
		return 0
	}

	return minLen
}