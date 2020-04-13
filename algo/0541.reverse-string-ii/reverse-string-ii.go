package problem0541

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	l := len(bytes)
	left := 0

	for i := 0; left < l; i++ {
		left = i * 2 * k
		right := left + k - 1

		if right > l-1 {
			right = l - 1
		}

		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}
