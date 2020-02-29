package problem0141

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
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
	one bool
}

func Test_Problem0141(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			para{
				[]int{3, 2, 0, -4},
				1,
			},
			ans{true},
		},
		question{
			para{
				[]int{1, 2},
				0,
			},
			ans{true},
		},
		question{
			para{
				[]int{1},
				-1,
			},
			ans{false},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		head := kit.Ints2ListWithCycle(p.one, p.two)

		//fmt.Printf("cycle link list %v \n", head)
		// head 类型转换
		//ast.Equal(a.one, hasCycle((*ListNode)(head)), "输入:%v", p)
		ast.Equal(a.one, hasCycle(head), "输入:%v \n", p)
	}
}
