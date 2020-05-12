package problem0215

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

var qs = []struct {
	para1 []int
	para2 int
	ans   int
}{
	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},
	{
		[]int{3, 2, 1, 5, 6, 4},
		1,
		6,
	},
	{
		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		4,
		4,
	},
}

func Test_FindKthLargest(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, findKthLargest(q.para1, q.para2), "输入:%v", q)
	}
}

func Benchmark_FindKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			findKthLargest(q.para1, q.para2)
		}
	}
}

func Test_FindKThLargeNumByQuick(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, findKThLargeNumByQuick(q.para1, q.para2), "输入:%v", q)
	}
}

func Benchmark_FindKThLargeNumByQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			findKThLargeNumByQuick(q.para1, q.para2)
		}
	}
}
