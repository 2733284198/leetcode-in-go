package problem0392

func isSubsequence(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	t1, t2 := 0, 0

	for t1 < len(s) && t2 < len(t) {
		if s[t1] == t[t2] {
			t1++
		}
		t2++
	}
	
	return t1 == len(s)
}