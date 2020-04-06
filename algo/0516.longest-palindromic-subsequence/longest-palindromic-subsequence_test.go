package problem0516

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  int
}{
	{
		"bbbab",
		4,
	},
	{
		"cbbd",
		2,
	},
}

func Test_LongestPalindromicSubsequence(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestPalindromicSubsequence(q.para), "输入：%v\n", q)
	}
}
