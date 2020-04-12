package problem0415

func addStrings(num1, num2 string) string {
	s1, s2 := len(num1), len(num2)

	if s1 < s2 { // convert s1 longest
		s1, s2, num1, num2 = s2, s1, num2, num1
	}

	result := make([]byte, s1+1)

	var overflow byte

	for i := s1 - 1; i >= 0; i-- {
		var a byte
		j := i - (s1 - s2)

		if j >= 0 {
			overflow, a = byteAdd(num1[i]+overflow, num2[j])
		} else {
			overflow, a = byteAdd(num1[i]+overflow, '0')
		}
		result[i+1] = a
	}

	if overflow > 0 {
		result[0] = '1'
		return string(result)
	}

	return string(result[1:])
}

func byteAdd(a, b byte) (byte, byte) {
	rtn := a + b - '0' - '0'

	if rtn >= 10 {
		return 1, rtn + '0' - 10
	}

	return 0, rtn + '0'
}
