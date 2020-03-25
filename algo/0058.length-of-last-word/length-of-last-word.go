package problem0058

import "strings"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	arr := strings.Fields(s)

	length := len(arr)

	if length > 0 {
		return len(arr[length-1])
	}

	return 0
}

func lengthOfLastWordUseLibV2(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	sArray := strings.Split(s, " ")

	return len(sArray[len(sArray)-1])
}

// lengthOfLastWordNotUseLib
func lengthOfLastWordNotUseLib(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		}
		if s[i] == ' ' && count != 0 {
			return count
		}
	}
	return count
}
