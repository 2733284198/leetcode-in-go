package problem0031

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := n - 2; i >= 0; i-- {
		for j := len(nums) - 1; j >= i; j-- {
			if i == 0 && j == 0 {
				sort.Ints(nums)
				return
			}
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				sort.Ints(nums[i+1:])
				return
			}
		}
	}
}
