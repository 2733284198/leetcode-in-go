package problem0080

func removeDuplicates(nums []int) int {
	n := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[n-1] {
			n++
			nums[n] = nums[i]
		}
	}

	return n + 1
}

func removeDuplicatesV2(nums []int) int {
	return removeDuplicatesN(nums, 2)
}

func removeDuplicatesN(nums []int, n int) int {
	fast, slow := n, n
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-n] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
