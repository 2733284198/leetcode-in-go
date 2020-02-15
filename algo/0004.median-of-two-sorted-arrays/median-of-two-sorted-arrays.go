package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return median(sliceMerge(nums1, nums2))
}

func median(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("nums length equal zero")
	}

	if l%2 == 0 {
		return (float64(nums[l/2-1]) + float64(nums[l/2])) / 2.0
	}

	return float64(nums[l/2])
}

func sliceMerge(a []int, b []int) []int {
	aLen, i := len(a), 0
	bLen, j := len(b), 0
	c := make([]int, aLen+bLen)
	for n := 0; n < aLen+bLen; n++ {
		if i == aLen || (i < aLen && j < bLen && a[i] > b[j]) {
			c[n] = b[j]
			j++
			continue
		}

		if j == bLen || (i < aLen && j < bLen && a[i] <= b[j]) {
			c[n] = a[i]
			i++
		}
	}

	return c
}

// a b is sorted array
func combineSort(a, b []int) []int {
	c := []int{}
	//i, j := 0
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	if i == len(a) {
		c = append(c, b[j:]...)
	}

	if j == len(b) {
		c = append(c, a[i:]...)
	}

	return c
}
