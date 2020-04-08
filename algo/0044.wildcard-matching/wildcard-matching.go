package problem0044

func isMatch(s string, p string) bool {
	dp := make([][]bool, 0, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]bool, len(p)+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i, _ := range s {
		dp[i+1][0] = false
	}
	for i, ch := range p {
		if dp[0][i] && ch == '*' {
			dp[0][i+1] = true
		} else {
			dp[0][i+1] = false
		}
	}
	for i, chS := range s {
		for j, chP := range p {
			x, y := i+1, j+1
			if chP != '*' {
				if dp[x-1][y-1] && (chP == '?' || chP == chS) {
					dp[x][y] = true
				} else {
					dp[x][y] = false
				}
			} else {
				if dp[x][y-1] || dp[x-1][y] {
					dp[x][y] = true
				} else {
					dp[x][y] = false
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
