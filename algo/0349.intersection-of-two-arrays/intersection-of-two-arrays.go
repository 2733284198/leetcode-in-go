package problem0349

func intersection(nums1 []int, nums2 []int) []int {

	hash := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range nums1 {
		hash[v] = true
	}

	for _, v := range nums2 {
		if hash[v] == true {
			result = append(result, v)
			hash[v] = false
		}
	}

	return result
}
