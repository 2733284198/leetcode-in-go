package problems0419

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var qs = []struct {
	para1 []int
	para2 int
	ans   []int
}{
	{
		[]int{2, 3, 12, 312, 4, 1, 5},
		3,
		[]int{1, 2, 3},
	},
	{
		[]int{1, 2, 3, 4},
		1,
		[]int{1},
	},
}

func Test_GetLeastNumbers(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, getLeastNumbers(q.para1, q.para2), "输入：%v\n", q)
	}
}
