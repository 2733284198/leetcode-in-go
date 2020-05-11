package problem0704

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		// mid := left + (right-left)/2
		mid := left + (right-left)>>1
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
