package problem0046

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
		[]int{1, 2, 3},
		[][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		},
	},
	//{
	//	[]int{1, 2, 8},
	//	[][]int{
	//		[]int{1, 2, 8},
	//		[]int{1, 8, 2},
	//		[]int{2, 1, 8},
	//		[]int{2, 8, 1},
	//		[]int{8, 1, 2},
	//		[]int{8, 2, 1},
	//	},
	//},
}

func Test_Permutations(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, permutations(q.para), "输入：%v\n", q)
	}
}

func Benchmark_Permutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			permutations(q.para)
		}
	}
}
