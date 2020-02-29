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
	three bool
}

var qs = []question{
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

func Test_Problem0141(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		head := kit.Ints2ListWithCycle(p.one, p.two)

		//fmt.Printf("cycle link list %v \n", head)
		// head 类型转换
		//ast.Equal(a.one, hasCycle((*ListNode)(head)), "输入:%v", p)
		ast.Equal(a.three, hasCycle(head), "输入:%v \n", p)
	}
}

func Benchmark_hasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, p := range qs {
			head := kit.Ints2ListWithCycle(p.one, p.two)
			hasCycle(head)
		}
	}
}