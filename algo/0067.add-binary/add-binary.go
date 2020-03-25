package problem0067

func addBinary(a string, b string) string {
	if len(a) < len(b) { // 左侧补齐
		a = padStart(a, len(b)-len(a))
	} else if len(b) < len(a) {
		b = padStart(b, len(a)-len(b))
	}

	res := make([]byte, len(a))

	borrow := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if borrow > 0 {
				res[i] = '1'
			} else {
				res[i] = '0'
				borrow++
			}

		} else if a[i] == '1' || b[i] == '1' {
			if borrow > 0 {
				res[i] = '0'
			} else {
				res[i] = '1'
			}
		} else {
			if borrow > 0 {
				res[i] = '1'
				borrow--
			} else {
				res[i] = '0'
			}
		}
	}
	if borrow > 0 {
		return "1" + string(res)
	}

	return string(res)
}

func padStart(s string, l int) string {
	start := make([]byte, l)
	for i := 0; i < len(start); i++ {
		start[i] = '0'
	}

	return string(start) + s
}
