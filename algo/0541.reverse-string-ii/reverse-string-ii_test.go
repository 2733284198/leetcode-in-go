package problem0541

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 string
	para2 int
	ans   string
}{
	{
		"abcdefg",
		2,
		"bacdfeg",
	},
}

func Test_ReverseString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, reverseStr(q.para1, q.para2), "输入：%v\n", q)
	}
}
