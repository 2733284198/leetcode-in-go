package problem0025

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		1,
		[]int{1, 2, 3, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		2,
		[]int{2, 1, 4, 3, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		3,
		[]int{3, 2, 1, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
		[]int{4, 3, 2, 1, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		5,
		[]int{5, 4, 3, 2, 1},
	},

	// more case
}

func Test_ReverseKGroup(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, kit.List2Ints(reverseKGroup(kit.Ints2List(q.para1), q.para2)), "输入：%v\n", q)
	}
}
