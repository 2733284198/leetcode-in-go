package problem0059

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para int
	ans  [][]int
}{
	{
		3,
		[][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
}

func Test_GenerateMatrix(t *testing.T) {

	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, generateMatrix(q.para), "输入:%v\n", q)
	}
}

func Benchmark_GenerateMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			generateMatrix(q.para)
		}
	}
}
