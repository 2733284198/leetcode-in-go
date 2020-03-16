package problem0088

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	para3 []int
	para4 int

	ans []int
}{
	{
		[]int{1, 2, 3, 0, 0, 0},
		3,
		[]int{2, 5, 6},
		3,
		[]int{1, 2, 2, 3, 5, 6},
	},
}

//func Test_MergeSortedArray(t *testing.T) {
//	ast := assert.New(t)
//
//	for _, q := range qs {
//		fmt.Printf("~~%v~~\n", q)
//		mergeSortedArray(q.para1, q.para2, q.para3, q.para4)
//		ast.Equal(q.ans, q.para1, "输入：%v", q)
//	}
//}

func Test_MergeSortedArrayNotUseBuiltInLib(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		mergeSortedArrayNotUseBuiltInLib(q.para1, q.para2, q.para3, q.para4)
		ast.Equal(q.ans, q.para1, "输入：%v", q)
	}

}

func Benchmark_MergeSortedArray(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
