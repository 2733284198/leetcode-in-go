package problem0028

import "strings"

func strStr(haystack string, needle string) int {

	// use built-in function

	return strings.Index(haystack, needle)
}

func strStrNotUseBuiltInFunc(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	needLen := len(needle)

	for i := 0; i <= len(haystack)-needLen; i++ {
		if haystack[i:i+needLen] == needle {
			return i
		}
	}

	return -1
}
