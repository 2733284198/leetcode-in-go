package prolbem0125

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !('a' <= s[l] && s[l] <= 'z' || '0' <= s[l] && s[l] <= '9') { // 左边 只处理字母和数字
			l++
			continue
		}
		if !('a' <= s[r] && s[r] <= 'z' || '0' <= s[r] && s[r] <= '9') {// 右边 只处理字母和数字
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
