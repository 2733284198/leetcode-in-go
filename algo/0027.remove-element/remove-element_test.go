package problem0027

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	val  int

	ans int
}{
	{
		[]int{3, 2, 2, 3},
		3,
		2,
	},
	{
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		5,
	},
}

func Test_RemoveElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeElement(q.para, q.val), "输入：%v", q)
	}
}

func Benchmark_RemoveElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			removeElement(q.para, q.val)
		}
	}
}
