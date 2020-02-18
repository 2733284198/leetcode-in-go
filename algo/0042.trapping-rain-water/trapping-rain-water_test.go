package problem0042

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
	one int
}

func Test_Problem0042(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{[]int{7, 7, 1, 7, 1, 0, 6, 9, 1, 6, 4, 9, 5, 8}},
			ans{39},
		},

		question{
			para{[]int{2, 0, 2}},
			ans{2},
		},

		question{
			para{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			ans{6},
		},

		question{
			para{[]int{2, 1}},
			ans{0},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, trap(p.one), "输入:%v", p)
	}
}
