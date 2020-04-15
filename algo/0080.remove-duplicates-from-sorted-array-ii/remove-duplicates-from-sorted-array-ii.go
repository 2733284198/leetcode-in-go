package problem0080

func removeDuplicates(nums []int) int {
	n := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[n - 1] {
			n ++
			nums[n] = nums[i]
		}
	}

	return n + 1
}