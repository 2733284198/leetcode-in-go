package problem0056

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
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		},
		[][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		},
	},
	{
		[][]int{
			{1, 4},
			{4, 5},
		},
		[][]int{
			{1, 5},
		},
	},
}

func Test_Merge(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, merge(q.para), "输入：%v\n", q)
	}
}

func Benchmark_Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
