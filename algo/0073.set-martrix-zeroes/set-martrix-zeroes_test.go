package problem0073

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]int
	ans  [][]int
}{
	{
		[][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		[][]int{
			[]int{1, 0, 1},
			[]int{0, 0, 0},
			[]int{1, 0, 1},
		},
	},
}

func Test_SetZeroes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		setZeroes(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}
