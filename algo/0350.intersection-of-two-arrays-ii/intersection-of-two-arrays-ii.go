package problem0350

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	result := make([]int, 0)

	for _, v := range nums1 {
		hash[v]++
	}

	for _, v := range nums2 {
		if hash[v] > 0 {
			result = append(result, v)
			hash[v]--
		}
	}

	return result
}
