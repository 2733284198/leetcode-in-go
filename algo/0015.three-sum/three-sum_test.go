package problem0015

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
		[]int{1, -1, -1, 0},
		[][]int{
			[]int{-1, 0, 1},
		},
	},
	{
		[]int{1, 2, -2},
		[][]int{
			[]int{-2, 1, 1},
		},
	},
	{
		[]int{0, 0, 0},
		[][]int{
			[]int{0, 0, 0},
		},
	},
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{
			[]int{-1, 0, 1},
			[]int{-1, -1, 2},
		},
	},
}

func Test_ThreeSum(t *testing.T) {
	ans := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ans.Equal(q.ans, threeSum(q.para), "输入：%v", q)
	}
}

func Benchmark_ThreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			threeSum(q.para)
		}
	}
}
