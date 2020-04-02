package problem0005

// longestPalindrome recursion
func longestPalindrome(s string) string {
	length := len(s)

	if length == 0 {
		return ""
	}

	max, maxL, maxR := 0, 0, 0

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			ret := isPalindrome(s, i, j)
			if ret && max < (j-i+1) {
				max = j - i + 1
				maxL = i
				maxR = j
			}
		}
	}

	// maxString := ""
	// for i := maxL; i <= maxR; i++ {
	// 	maxString += string(s[i])
	// }

	// return maxString

	return s[maxL : maxR+1]

}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}

// 第二种思路：遍历字符串，找出以每个字符为中心的回文字符串有多长，选最长的返回
func longestPalindromeByDp(s string) string {
	length := len(s)
	getLen := func(i, j int) int {
		// 以s[i]s[j]为中心的最长回文字符串
		for i >= 0 && j < length {
			if s[i] == s[j] {
				i--
				j++
			} else {
				return j - i - 1
			}
		}
		return j - i - 1
	}
	max := 0
	maxStart := 0
	for i := 0; i < length; i++ {
		if Max(getLen(i, i+1), getLen(i, i)) > max {
			max = Max(getLen(i, i+1), getLen(i, i))
			maxStart = i - (max-1)/2
		}
	}
	maxString := ""
	for i := maxStart; i < maxStart+max; i++ {
		maxString += string(s[i])
	}
	return maxString
}

func Max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
