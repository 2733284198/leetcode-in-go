package problem0115

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 string
	ans   int
}{
	{
		"rabbbit",
		"rabbit",
		3,
	},
	{
		"babgbag",
		"bag",
		5,
	},
}

func Test_NumDistinct(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, numDistinct(q.para1, q.para2), "输入：%v\n", q)
	}
}
