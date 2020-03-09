package problem0088

import "sort"

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1) // 依赖标准库
}
