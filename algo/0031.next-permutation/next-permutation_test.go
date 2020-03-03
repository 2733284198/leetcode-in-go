package problem0031

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},
	{
		[]int{3, 2, 1},
		[]int{1, 2, 3},
	},
	{
		[]int{1, 1, 5},
		[]int{1, 5, 1},
	},
}

func Test_NextPermutation(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, nextPermutation(q.para), "输入：%v", q)
	}
}

func Benchmark_NextPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			nextPermutation(q.para)
		}
	}
}
