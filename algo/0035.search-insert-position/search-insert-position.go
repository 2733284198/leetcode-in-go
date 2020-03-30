package problem0035

func searchInsert(nums []int, target int) int {
	index := 0
	for _, num := range nums {
		if num < target {
			index += 1
		}
	}

	return index
}
