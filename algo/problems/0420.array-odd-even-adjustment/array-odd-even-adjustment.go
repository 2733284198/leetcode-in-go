package problems0420

func arrayAdjust(nums []int)  {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	oddCount, evenCount := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 1 {
			oddCount ++
		}
	}

	for i := 0; i < len(nums2); i++ {
		if nums2[i]%2 == 0 {
			 nums[oddCount] = nums2[i]
			 oddCount ++
		} else {
			nums[evenCount] = nums2[i]
			evenCount ++
		}
	}
}

func arrayAdjustV2(nums []int) []int {
	ret := make([]int, len(nums))
	start, end := 0, len(nums) -1

	for _,v := range nums {
		if v %2 == 1 {
			ret[start] = v
			start ++
		} else {
			ret[end] = v
			end --
		}
	}

	return ret
}