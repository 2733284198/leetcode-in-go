package problem0063

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
			[]int{0, 0, 0},
			[]int{0, 1, 0},
			[]int{0, 0, 0},
		},
		2,
	},
}

func Test_UniquePaths(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, uniquePathsWithObstacles(q.para), "输入：%v", q)
	}
}
