package problem0053

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
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
	{
		[]int{},
		0,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		55,
	},
	{
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		-1,
	},
	{
		[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, 3},
		3,
	},
}

func Test_MaxSubArray(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxSubArray(q.para), "输入：%v\n", q)
	}
}

func Benchmark_MaxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			maxSubArray(q.para)
		}
	}
}
