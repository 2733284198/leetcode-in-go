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
	// maxString := ""
	// for i := maxStart; i < maxStart+max; i++ {
	// 	maxString += string(s[i])
	// }
	// return maxString

	return s[maxStart : maxStart+max]
}

func Max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func longestPalindromeByDpV2(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	start, max := 0, 1

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (dp[j+1][i-1] || i-j <= 2) {
				dp[j][i] = true
			} else {
				dp[j][i] = false
			}

			if dp[j][i] {
				cur := i - j + 1
				if cur > max {
					max = cur
					start = j
				}
			}
		}
	}

	return s[start : start+max]
}

func longestPalindromeV3(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}

		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		// 为下一次循环做准备
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}
