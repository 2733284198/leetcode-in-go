package problem0020

func isValid(s string) bool {
	if s == "" {
		return true
	}

	l := len(s)
	if l%2 == 1 {
		return false
	}

	var stack []byte

	for i := 0; i < l; i++ {

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		}

		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 {
				return false
			}
			topElement := stack[len(stack)-1]
			if isPair(topElement, s[i]) == false {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return 0 == len(stack)
}

func isPair(ch1, ch2 byte) bool {
	if ch1 == '(' && ch2 == ')' ||
		ch1 == '[' && ch2 == ']' ||
		ch1 == '{' && ch2 == '}' {
		return true
	}
	return false
}
