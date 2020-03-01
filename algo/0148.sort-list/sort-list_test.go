package problem0148

import (
	"fmt"
	"testing"

	"github.com/aierui/leetcode-in-go/kit"
	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	head []int
	ans  []int
}{
	{
		[]int{4, 2, 1, 3},
		[]int{1, 2, 3, 4},
	},
	{
		[]int{-1, 5, 3, 4, 0},
		[]int{-1, 0, 3, 4, 5},
	},
}

func Test_SortList(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		head := kit.Ints2List(q.head)
		ast.Equal(q.ans, kit.List2Ints(sortList(head)), "输入：%v\n", q)
	}
}

func Benchmark_SortList(b *testing.B) {
	head := kit.Ints2List([]int{9, 6, 3, 1, 4, 8, 2, 5, 7, 19, 16, 13, 11, 14, 18, 12, 15, 17, 29, 26, 23, 21, 24, 28, 22, 25, 27, 39, 36, 33, 31, 34, 38, 32, 35, 37, 49, 46, 43, 41, 44, 48, 42, 45, 47})
	for i := 0; i < b.N; i++ {
		sortList(head)
	}
}
