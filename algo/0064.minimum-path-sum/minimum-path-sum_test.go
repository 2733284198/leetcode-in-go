package problem0064

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]int
	ans  int
}{
	{
		[][]int{
			[]int{1, 3, 1},
			[]int{1, 5, 1},
			[]int{4, 1, 1},
		},
		7,
	},
}

func Test_MiniMumPathSum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, minPathSum(q.para), "输入：%v\n", q)
	}
}
