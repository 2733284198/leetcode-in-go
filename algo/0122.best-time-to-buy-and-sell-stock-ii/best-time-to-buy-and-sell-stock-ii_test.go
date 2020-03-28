package problem0122

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  int
}{
	{
		[]int{1, 2, 3, 4, 5},
		4,
	},
	{
		[]int{7, 6, 5, 4, 3, 2},
		0,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{7, 2, 5, 3, 2, 1},
		3,
	},
	{
		[]int{7, 3, 5, 3, 1, 2},
		2,
	},
	{
		[]int{7, 1, 5, 7, 2, 2},
		6,
	},
}

func Test_MaxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxProfit(q.para), "输入：%v\n", q)
	}
}
