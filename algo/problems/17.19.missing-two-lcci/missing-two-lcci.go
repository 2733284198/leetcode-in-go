package problems1719

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	S0 := ((1 + n) * n) / 2
	S1, a, b, mid := 0, 0, 0, 0
	for _, num := range nums {
		S1 += num
	}
	mid = (S0 - S1) / 2
	midTotal := ((1 + mid) * mid) / 2
	for _, num := range nums {
		if num <= mid {
			midTotal = midTotal - num
		}
	}

	a = midTotal
	b = S0 - S1 - a

	return []int{a, b}
}
