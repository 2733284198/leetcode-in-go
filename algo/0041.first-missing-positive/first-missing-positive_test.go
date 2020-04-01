package problem0041

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
		[]int{1, 2, 0},
		3,
	},
	{
		[]int{3, 4, -1, -1},
		2,
	},
	{
		[]int{7, 8, 9, 11, 12},
		1,
	},
}

func Test_FirstMissingPositive(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~", q)
		ast.Equal(q.ans, firstMissingPositive(q.para), "输入：%v", q)
	}
}
