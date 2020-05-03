package problem0300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para []int
	ans  int
}{
	{
		[]int{10, 9, 2, 5, 3, 7, 101, 18},
		4,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		8,
	},
}

func Test_LengthOfLIS(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLIS(q.para), "输入：%v\n", q)
	}
}

func Benchmark_LengthOfLIS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			lengthOfLIS(q.para)
		}
	}
}

func Test_LengthOfLISByBinarySearch(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, lengthOfLISByBinarySearch(q.para), "输入：%v\n", q)
	}
}

func Benchmark_LengthOfLISByBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			lengthOfLISByBinarySearch(q.para)
		}
	}
}
