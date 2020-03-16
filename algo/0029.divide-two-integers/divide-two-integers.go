package problem0029

import "math"

func divide(dividend, divisor int) int {
	isNegative := (dividend >> 31 & 1) ^ (divisor >> 31 & 1)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	var m int64 // 商

	base := divisor
	for base<<1 <= dividend {
		base <<= 1 // base 左移几次，就表示商 m 的二进制有几位
	}

	for base >= divisor {
		t := 0
		if base <= dividend {
			t = 1
		}
		m = m<<1 + int64(t)
		dividend -= base * t
		base >>= 1 // base 每次右移一位，当它小于初值时，就求得了商 m 的所有位
	}

	if isNegative == 1 {
		m = -m
	}
	if m < math.MinInt32 || m > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(m)
}
