package kit

func Bubble(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func Insert(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := nums[i]
		j := i - 1
		/**
		 * 当已排序部分的当前元素大于value，
		 * 就将当前元素向后移一位，再将前一位与value比较
		 */
		for ; j >= 0 && nums[j] > value; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = value
	}
}

func Selection(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		// 查找最小值
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 交换
		nums[i], nums[min] = nums[min], nums[i]
	}
}
