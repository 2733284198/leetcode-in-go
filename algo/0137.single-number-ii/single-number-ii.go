package problem0137

func singleNumberHashMap(nums []int) int {
	m := make(map[int]int)

	for _, k := range nums {
		m[k]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}

func singleNumber(nums []int) int {
	i, j := 0, 0

	for _, v := range nums {
		i = (i ^ v) & ^j
		j = (j ^ v) & ^i
	}

	return i
}
