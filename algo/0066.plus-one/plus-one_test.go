package problem0066

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 2, 4},
	},
	{
		[]int{4, 3, 2, 1},
		[]int{4, 3, 2, 2},
	},
	{
		[]int{9},
		[]int{1, 0},
	},
	{
		[]int{9, 9},
		[]int{1, 0, 0},
	},
}

func Test_PlusOne(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, plusOne(q.para), "输入：%v\n", q)
	}
}
