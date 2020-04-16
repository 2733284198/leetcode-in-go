package problem0026

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func removeDuplicatesV2(nums []int) int {
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[n] != nums[i] {
			n++
			nums[n] = nums[i]
		}
	}
	return n + 1
}

func removeDuplicatesV3(nums []int) int {
	return removeDuplicatesN(nums, 1)
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
