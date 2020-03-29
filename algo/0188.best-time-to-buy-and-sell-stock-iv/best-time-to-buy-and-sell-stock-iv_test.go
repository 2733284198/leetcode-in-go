package problem0188

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   int
}{
	{
		[]int{1, 2, 3, 4, 5},
		2,
		4,
	},
	{
		[]int{7, 6, 5, 4, 3, 2},
		3,
		0,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		4,
		7,
	},
	{
		[]int{7, 2, 5, 3, 2, 1},
		1,
		3,
	},
	{
		[]int{7, 3, 5, 3, 1, 2},
		3,
		3,
	},
	{
		[]int{7, 1, 5, 7, 2, 2},
		4,
		6,
	},
	{
		[]int{2, 1, 5, 7, 1, 9, 1, 1, 31, 10, 100, 1, 19},
		6,
		152,
	},
}

func Test_MaxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxProfit(q.para2, q.para1), "输入：%v\n", q)
	}
}
