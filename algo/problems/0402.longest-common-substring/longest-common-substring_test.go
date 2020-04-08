package problems0402

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 string
	ans          string
}{
	{
		"abcefs", "abc",
		"abc",
	},
	{
		"ddaabcefs", "abc",
		"abc",
	},
	{
		"ddaabcefs", "abcd",
		"abc",
	},
	{
		"ddaabcefs", "labcd",
		"abc",
	},
	{
		"ddaabcefs", "d",
		"d",
	},
	{
		"ddaabcefs", "z",
		"",
	},
}

func Test_LCS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxSubTwoString(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Test_LCSV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestCommonSubString(q.para1, q.para2), "输入：%v\n", q)
	}
}
