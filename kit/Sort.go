package kit

/*
* 冒泡排序
*
* 时间复杂度 平均 最好 最坏 O(n^2) O(n) O(n^2)
* 空间复杂度 O(1)
* 稳定
 */
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

/*
* 插入排序
*
* 时间复杂度 平均 最好 最坏 O(n^2) O(n) O(n^2)
* 空间复杂度 O(1)
* 稳定
 */
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

/*
* 选择排序
*
* 时间复杂度 平均 最好 最坏 O(n2) O(n2) O(n2)
* 空间复杂度 O(1)
* 不稳定
 */
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

/*
* 快速排序
* 时间复杂度 平均 最好 最坏 O(nlog(n)) O(nlog(n)) O(n^2)
* 空间复杂度 O(nlog(n))
* 不稳定
*
 */
func Quick(nums []int) {
	n := len(nums)

	if n <= 1 {
		return
	}

	quick(nums, 0, n-1)
}

func quick(nums []int, start, end int) {
	if end < start {
		return
	}

	pivot := nums[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			if splitIndex != i {
				nums[splitIndex], nums[i] = nums[i], nums[splitIndex]
			}
			splitIndex++
		}
	}

	nums[end], nums[splitIndex] = nums[splitIndex], pivot

	quick(nums, start, splitIndex-1)
	quick(nums, splitIndex+1, end)
}

/*
* 归并排序
* 时间复杂度 平均 最好 最坏 O(nlog(n)) O(nlog(n)) O(nlog(n))
* 空间复杂度 O(n)
* 稳定
*
 */
func Merge(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	mergeRecursion(nums, 0, n-1)
}

func mergeRecursion(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeRecursion(nums, start, mid)
	mergeRecursion(nums, mid+1, end)

	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if nums[i] <= nums[j] {
			tmpArr[k] = nums[i]
			i++
		} else {
			tmpArr[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = nums[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = nums[j]
		k++
	}
	copy(nums[start:end+1], tmpArr)
}

/**
* 堆排序
*
* 时间复杂度 平均 最好 最坏 O(nlog(n)) O(nlog(n)) O(nlog(n))
* 空间复杂度 O(1)
*
* 不稳定
*
 */
func Heap(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	h := (n - 1) >> 1
	for i := h; i >= 0; i-- {
		heap(nums, i, n-1)
	}

	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heap(nums, 0, i-1)
	}
}

func heap(nums []int, start, end int) {
	left := start*2 + 1
	right := left + 1

	if left > end {
		return
	}

	tmp := left
	if right <= end && nums[right] > nums[left] {
		tmp = right
	}

	if nums[tmp] > nums[start] {
		nums[start], nums[tmp] = nums[tmp], nums[start]
		heap(nums, tmp, end)
	}
}
