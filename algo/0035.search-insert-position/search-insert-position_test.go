package problem0035

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
		[]int{1, 3, 5, 6},
		5,
		2,
	},
	{
		[]int{1, 3, 5, 6},
		2,
		1,
	},
	{
		[]int{1, 3, 5, 6},
		7,
		4,
	},
	{
		[]int{1, 3, 5, 6},
		0,
		0,
	},
}

func Test_SearchInsert(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		fmt.Printf("~~%v~~\n", q)
		ast.Equal(q.ans, searchInsert(q.para1, q.para2), "输入：%v\n", q)
	}
}
