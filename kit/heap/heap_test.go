package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var qs = []struct {
	para []int
	ans  []int
}{
	{
		[]int{2, 3, 32, 3, 4, 1, 7, 10, 66},
		[]int{1, 2, 3, 3, 4, 7, 10, 32,66},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func Test_HeapSort(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		Sort(q.para)
		ast.Equal(q.ans, q.para, "输入：%v\n", q)
	}
}

func Benchmark_HeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			Sort(q.para)
		}
	}
}
