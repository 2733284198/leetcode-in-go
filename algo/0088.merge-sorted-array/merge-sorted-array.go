package problem0088

import "sort"

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1) // 依赖标准库
}

// mergeSortedArrayNotUseBuiltInLib
// not need to use extra space
// in place merge
func mergeSortedArrayNotUseBuiltInLib(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1  // pointer nums1 max(last) value
	second := n - 1 // pointer nums2 max(last) value

	// m + n - 1 => pointer num1 end value
	for i := m + n - 1; i >= 0; i-- {
		if second < 0 { // nums2 has already been merged.
			break
		}

		if first >= 0 && nums1[first] > nums2[second] {
			nums1[i] = nums1[first]
			first--
		} else {
			nums1[i] = nums2[second]
			second--
		}
	}
}
