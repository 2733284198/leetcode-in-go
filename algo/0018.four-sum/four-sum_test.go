package problem0018

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
		[]int{1, 0, -1, 0, -2, 2},
		[][]int{
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
			[]int{-1, 0, 0, 1},
		},
	},
}

func Test_FourSum(t *testing.T) {
	ans := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ans.Equal(q.ans, fourSum(q.para, 0), "输入：%v", q)
	}
}

func Benchmark_FourSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			fourSum(q.para, 0)
		}
	}
}
