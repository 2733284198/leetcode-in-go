package problem0005

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  string
}{

	{
		"babad",
		"bab",
	},
	{
		"cbbd",
		"bb",
	},

	{
		"cbbdddadaddaadggaedewe",
		"ddadadd",
	},
}

func Test_LongestPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestPalindrome(q.para), "输入：%v\n", q)
	}
}

func Test_LongestPalindromeByDp(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestPalindromeByDp(q.para), "输入：%v\n", q)
	}
}
