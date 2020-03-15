package problem0047

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  [][]int
}{
	{
		[]int{1, 1, 2},
		[][]int{
			[]int{1, 1, 2},
			[]int{1, 2, 1},
			[]int{2, 1, 1},
		},
	},
}

func Test_PermuteUnique(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, permuteUnique(q.para), "输入：%v\n", q)
	}
}

func Test_PermuteUniqueV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, permuteUniqueV2(q.para), "输入：%v\n", q)
	}
}
