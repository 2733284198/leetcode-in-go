package problem1143

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 string
	ans          int
}{
	{
		"abcde",
		"ace",
		3,
	},
	{
		"abc",
		"abc",
		3,
	},
	{
		"abc",
		"def",
		0,
	},
}

func Test_LongestCommonSubsequence(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestCommonSubsequence(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Test_LongestCommonSubsequenceV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestCommonSubsequenceV2(q.para1, q.para2), "输入：%v\n", q)
	}
}
