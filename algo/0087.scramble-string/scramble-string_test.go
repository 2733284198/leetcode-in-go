package problem0087

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
		"great",
		"rgeat",
		true,
	},
	{
		"abcde",
		"caebd",
		false,
	},
}

func Test_IsScramble(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isScramble(q.para1, q.para2), "输入：%v\n", q)
	}
}
