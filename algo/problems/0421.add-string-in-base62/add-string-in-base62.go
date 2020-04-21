package problems0421

import (
	"strconv"
	"strings"
)

func addStringInBase62(a, b string) string {
	aLen, bLen := len(a), len(b)
	if aLen > bLen {
		b = strings.Repeat("0", aLen-bLen) + b
	} else {
		a = strings.Repeat("0", bLen-aLen) + a
	}

	var result []string
	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		sum := int(getValue(rune(a[i]))) + int(getValue(rune(b[i]))) + carry

		if sum >= 62 {
			carry = 1
		}

		rtn := toChar(sum % 62)
		result = append(result, rtn)
	}

	if carry == 1 {
		result = append(result, "1")
	}

	var ret string

	for _, v := range result {
		ret = v + ret
	}

	return ret
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
