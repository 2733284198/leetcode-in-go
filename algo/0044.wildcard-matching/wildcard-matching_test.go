package problem0044

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   bool
}{
	{
		"aa",
		"a",
		false,
	},
	{
		"aa",
		"*",
		true,
	},
	{
		"cb",
		"?a",
		false,
	},
	{
		"adceb",
		"*a*b",
		true,
	},
	{
		"acdcb",
		"a*c?b",
		false,
	},
}

func Test_IsMatch(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isMatch(q.para1, q.para2), "输入：%v\n", q)
	}
}
