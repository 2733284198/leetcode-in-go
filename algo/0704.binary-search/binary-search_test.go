package problem0704

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		4,
		3,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		10,
		-1,
	},
}

func Test_Search(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, search(q.para1, q.para2), "输入：%v\n", q)
	}
}

func Benchmark_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			search(q.para1, q.para2)
		}
	}
}
