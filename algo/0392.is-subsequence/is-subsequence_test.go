package problem0392

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1, para2 string
	ans          bool
}{
	{
		"abc",
		"ahbgdc",
		true,
	},
	{
		"axc",
		"ahbgdc",
		false,
	},
}

func Test_IsSubsequence(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isSubsequence(q.para1, q.para2), "输入：%v\n", q)
	}
}
