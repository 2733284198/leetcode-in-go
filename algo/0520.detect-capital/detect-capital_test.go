package problem0520

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para string
	ans  bool
}{
	{
		"ABC",
		true,
	},
	{
		"Abc",
		true,
	},
	{

		"abc",
		true,
	},
	{
		"ABBc",
		false,
	},
}

func Test_DetectCapitalUse(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, detectCapitalUse(q.para), "输入：%v\n", q)
	}
}
