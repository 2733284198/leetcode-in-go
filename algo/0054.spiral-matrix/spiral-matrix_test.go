package problem0054

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para [][]int
	ans  []int
}{
	{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
}

func Test_SpiralOrder(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, spiralOrder(q.para), "输入：%v\n", q)
	}
}

func Benchmark_SpiralOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			spiralOrder(q.para)
		}
	}
}
