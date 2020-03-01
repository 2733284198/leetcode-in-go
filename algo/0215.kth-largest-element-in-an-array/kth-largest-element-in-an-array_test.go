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

var qs = []question{
	question{
		para{
			[]int{3, 2, 1, 5, 6, 4},
			2,
		},
		ans{
			5,
		},
	},
	question{
		para{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
		},
		ans{
			4,
		},
	},
}

func Test_Problem0215(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, findKthLargest(p.one, p.two), "输入:%v", p)
	}
}

func Benchmark_Problem0215(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			_, p := q.ans, q.para
			findKthLargest(p.one, p.two)
		}
	}
}
