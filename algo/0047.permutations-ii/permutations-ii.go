package problem0047

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backTrace(0, nums, &res)
	return res
}

func backTrace(first int, nums []int, res *[][]int) {
	if first == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}
	m := map[int]bool{}
	for i := first; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[first], nums[i] = nums[i], nums[first]
		backTrace(first+1, nums, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}
