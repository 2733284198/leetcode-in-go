package problem0123

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
		[]int{1, 2, 3, 4, 5},
		4,
	},
	{
		[]int{7, 6, 5, 4, 3, 2},
		0,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		7,
	},
	{
		[]int{7, 2, 5, 3, 2, 1},
		3,
	},
	{
		[]int{7, 3, 5, 3, 1, 2},
		3,
	},
	{
		[]int{7, 1, 5, 7, 2, 2},
		6,
	},
	{
		[]int{2, 1, 5, 7, 1, 9},
		14,
	},
}

func Test_MaxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxProfit(q.para), "输入：%v\n", q)
	}
}

func Test_MaxProfitV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, maxProfitV2(q.para), "输入：%v\n", q)
	}
}

// func Test_MaxProfitV3(t *testing.T) {
// 	ast := assert.New(t)
// 	for _, q := range qs {
// 		fmt.Printf("~~%v~~\n", q)
// 		ast.Equal(q.ans, maxProfitV3(q.para), "输入：%v\n", q)
// 	}
// }

func Benchmark_MaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			maxProfit(q.para)
		}
	}
}

func Benchmark_MaxProfitV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			maxProfitV2(q.para)
		}
	}
}
