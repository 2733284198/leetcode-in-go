package problem0142

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	ints []int
	pos  int
}{
	{
		[]int{1},
		-1,
	},
	{
		[]int{1, 2, 3},
		-1,
	},
	{
		[]int{3, 2, 0, -4},
		1,
	},
	{
		[]int{1, 2},
		0,
	},
}

func Test_Problem0142(t *testing.T) {
	ast := assert.New(t)
	for _, p := range qs {
		fmt.Printf("~~%v~~\n", p)
		head := kit.Ints2ListWithCycle(p.ints, p.pos)
		var ans *ListNode
		if p.pos >= 0 {
			ans = head.GetNodeWith(p.ints[p.pos])
		}
		ast.Equal(ans, detectCycle(head), "输入:%v", p)
	}
}

func Benchmark_Problem0142(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
