package problem0022

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  []string
}{
	{
		3,
		[]string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		},
	},
}

func Test_GengerateParenthesis(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, generateParenthesis(q.para), "输入：%v\n", q)
	}
}
