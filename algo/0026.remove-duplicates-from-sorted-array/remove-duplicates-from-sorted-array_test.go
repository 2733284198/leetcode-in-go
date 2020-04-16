package problem0026

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
		[]int{1, 1, 2},
		2,
	},
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
}

func Test_RemoveDuplicates(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeDuplicates(q.para), "输入：%v\n", q)
	}
}

func Test_RemoveDuplicatesV2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeDuplicatesV2(q.para), "输入：%v\n", q)
	}
}
func Test_RemoveDuplicatesV3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, removeDuplicatesV3(q.para), "输入：%v\n", q)
	}
}

func Benchmark_RemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			removeDuplicates(q.para)
		}
	}
}
