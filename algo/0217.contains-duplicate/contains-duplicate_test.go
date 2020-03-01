package problem0217

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
}

type ans struct {
	one bool
}

var qs = []question{
	question{
		para{
			[]int{1, 2, 3, 1},
		},
		ans{
			true,
		},
	},
	question{
		para{
			[]int{1, 2, 3, 4},
		},
		ans{
			false,
		},
	},
	question{
		para{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		},
		ans{
			true,
		},
	},
}

func Test_Problem0217(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, containsDuplicate(p.one), "输入:%v", p)
	}
}

func Benchmark_Problem0217(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			_, p := q.ans, q.para
			containsDuplicate(p.one)
		}
	}
}
