package kit

// BinarySearch
func BinarySearch(nums []int, val int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)>>1

		if nums[mid] > val {
			end = mid - 1
		} else if nums[mid] < val {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
