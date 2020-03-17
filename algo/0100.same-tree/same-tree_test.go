package problem0100

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 []int
	para3 []int
	para4 []int
	ans   bool
}{
	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		false,
	},
}

func Test_IsSameTree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, isSameTree(kit.PreIn2Tree(q.para1, q.para2), kit.PreIn2Tree(q.para3, q.para4)), "输入：%v\n", q)
	}
}

func Benchmark_IsSameTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			isSameTree(kit.PreIn2Tree(q.para1, q.para2), kit.PreIn2Tree(q.para3, q.para4))
		}
	}
}
