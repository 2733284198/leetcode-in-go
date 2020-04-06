package problem0032

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
		"(()",
		2,
	},
	{
		")()())",
		4,
	},
}

func Test_LongestValidParentheses(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestValidParentheses(q.para), "输入：%v\n", q)
	}
}

func Test_LongestValidParenthesesByStack(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, longestValidParenthesesByStack(q.para), "输入：%v\n", q)
	}
}
