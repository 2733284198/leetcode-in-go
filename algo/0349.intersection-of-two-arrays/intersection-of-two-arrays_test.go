package problem0349

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	ans   []int
}{
	{
		[]int{1, 2, 2, 1},
		[]int{2, 2},
		[]int{2},
	},
	{
		[]int{4, 9, 5},
		[]int{9, 4, 9, 8, 4},
		[]int{9, 4},
	},
}

func Test_Intersection(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, intersection(q.para1, q.para2), "输入：%v\n", q)
	}
}
