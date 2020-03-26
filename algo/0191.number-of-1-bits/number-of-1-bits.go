package problem0191

func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		n++
		num = num & (num - 1)
	}
	return n
}
