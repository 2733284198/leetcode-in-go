package problem0057

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 [][]int
	para2 []int
	ans   [][]int
}{
	{
		[][]int{
			{1, 3},
			{6, 9},
		},
		[]int{2, 5},
		[][]int{
			{1, 5},
			{6, 9},
		},
	},
	{
		[][]int{
			{1, 2},
			{3, 5},
			{6, 7},
			{8, 10},
			{12, 16},
		},
		[]int{4, 8},
		[][]int{
			{1, 2},
			{3, 10},
			{12, 16},
		},
	},
}

func Test_Merge(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, insert(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Benchmark_Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
