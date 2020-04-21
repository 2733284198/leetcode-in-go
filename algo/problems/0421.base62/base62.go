package problems0421

import (
	"strconv"
)

func base62To10(s string) int {
	var r int
	for _, v := range []byte(s) {
		i := int(getValue(rune(v)))
		r = 62*r + int(i)
	}
	return r
}

func base10To62(num int) string {
	str := make([]byte, 0)
	for {
		if num <= 0 {
			break
		}
		i := num % 62 // 余数
		str = append([]byte(toChar(i)), str...)
		num /= 62 // 商
	}

	return string(str)
}

func toChar(i int) string {
	if i >= 0 && i <= 9 {
		return strconv.Itoa(i)
	}

	if i > 9 && i < 36 {
		return string('a' + i - 10)
	}

	return string('A' + i - 36)
}

func getValue(s rune) rune {
	if s >= '0' && s <= '9' {
		return s - '0'
	}

	if s >= 'a' && s <= 'z' {
		return s - 'a' + 10
	}

	return s - 'A' + 36
}
