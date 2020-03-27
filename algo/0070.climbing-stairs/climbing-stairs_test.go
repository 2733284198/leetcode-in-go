package problem0070

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  int
}{
	{
		2,
		2,
	},
	{
		3,
		3,
	},
}

func Test_ClimbStairs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, climbStairs(q.para), "输入：%v\n", q)
	}
}
